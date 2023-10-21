package error_handler_test

import (
	"log"
	"testing"
	errorhandler "url-shortener/pkg/error"
)

func TestHandleError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "Test HandleError with AppError",
			err:      errorhandler.NewAppError(500, "Internal Server Error"),
			expected: "Error: Internal Server Error\n",
		},
		{
			name:     "Test HandleError with nil error",
			err:      nil,
			expected: "", // No output expected for nil error
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Redirect the log output to a buffer for testing
			logOutput := captureLogOutput(func() {
				errorhandler.HandleError(test.err)
			})

			if logOutput != test.expected {
				t.Errorf("Expected log output: %s, got: %s", test.expected, logOutput)
			}
		})
	}
}

func captureLogOutput(fn func()) string {
	// Capture log output to a buffer for testing
	var logOutput string
	log.SetOutput(&testLogWriter{&logOutput})
	defer log.SetOutput(log.Writer())
	fn()
	return logOutput
}

type testLogWriter struct {
	output *string
}

func (tlw *testLogWriter) Write(p []byte) (n int, err error) {
	*tlw.output = string(p)
	return len(p), nil
}
