package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"url-shortener/api/handlers"
	_ "url-shortener/docs"
	cacheManagement "url-shortener/pkg/cache"
	configManagement "url-shortener/pkg/config-reader"
	logManagement "url-shortener/pkg/log"
)

// @title			Swagger Example API
// @version		1.0
// @description	This is a sample server celler server.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host			localhost:9191
// @BasePath		/api/v1
func main() {
	path := "C:/Users/rafiee/Desktop/WS-GO/src/go-url-shortener/configuration/config.yml"
	configManagement.ConfigFactory(path)
	logManagement.LoggerFactory()
	cacheManagement.ManagementFactory()
	engine := gin.Default()
	defer func() {
		config := configManagement.Instance()
		address := config.Server.Domain + ":" + config.Server.Port
		err := engine.Run(address)
		if err != nil {
			return
		}
	}()
	handlers.NewShorterConfigRoutes(engine)
	configSwagger(engine)
	configCORS(engine)
}

func configCORS(engine *gin.Engine) {
	// Create a CORS configuration with your desired settings.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	engine.Use(cors.New(corsConfig))
}

func configSwagger(engine *gin.Engine) {
	// Serve the Swagger documentation on /swagger URL
	// http://url/swagger/index.html
	engine.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
}
