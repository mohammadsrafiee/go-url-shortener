package errorHandler

import (
	"log"
)

// AppError represents a custom application error.
type AppError struct {
	Code    int    // HTTP status code
	Message string // Error message
}

// NewAppError creates a new AppError.
func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Error returns the error message
func (e *AppError) Error() string {
	return e.Message
}

// HandleError is a function to handle errors and log them.
func HandleError(err error) {
	if err == nil {
		return
	}
	// Log the error
	log.Printf("Error: %v\n", err)
	// You can also implement additional error handling logic here, such as sending error responses to clients.
}

// Usage:
// In your application code, when an error occurs, create an AppError and pass it to the HandleError function.

func ExampleUsage() {
	// Simulate an error
	err := NewAppError(500, "Internal Server Error")

	// Handle the error
	HandleError(err)
}
