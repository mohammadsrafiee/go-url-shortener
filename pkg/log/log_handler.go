package logHandler

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
	"time"
	configReader "url-shortener/pkg/config-reader"
)

var (
	logger *zap.Logger
)

func Logger() *zap.Logger {
	return logger
}

func LogWriter() zapcore.WriteSyncer {
	var config = configReader.Instance()
	logRotate, _ := rotatelogs.New(
		config.Log.Path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(config.Log.Path),
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 7 days
		rotatelogs.WithRotationTime(24*time.Hour), // 24 hours
	)
	return zapcore.AddSync(logRotate)
}

func LoggerFactory() {
	var config = configReader.Instance()
	level, _ := zapcore.ParseLevel(string(rune(config.Log.Level)))
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(LogWriter()),
		level,
	), zap.AddCaller())
	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {
			// nothing to do
		}
	}(logger)
}

// LoggingMiddleware is a middleware function that logs function entry and exit.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		Logger().Debug("Entered:  " + r.Method + " " + r.URL.Path)
		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
		// Calculate the elapsed time.
		elapsed := time.Since(startTime)
		Logger().Debug("Exited: " + r.Method + " " + r.URL.Path + ", Elapsed Time: %s" + strconv.FormatInt(int64(elapsed), 10))
	})
}
