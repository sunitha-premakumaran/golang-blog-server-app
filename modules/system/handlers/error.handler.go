package handlers

type AppError struct {
	Message       string
	CorrelationId string
	StatusCode    int
	Details       interface{}
}

// TODO: add correlationId logic
func NewHTTPError(statusCode int, message string, details interface{}) *AppError {
	error := AppError{Message: message, StatusCode: statusCode, Details: details}
	return &error
}
