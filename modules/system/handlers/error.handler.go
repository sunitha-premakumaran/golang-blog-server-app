package handlers

type AppError struct {
	Message       string      `json:"message"`
	CorrelationId string      `json:"correlationId"`
	StatusCode    int         `json:"statusCode"`
	Details       interface{} `json:"details"`
}

func (m *AppError) Error() string {
	return "error"
}

// TODO: add correlationId logic
func NewHTTPError(statusCode int, message string, details interface{}) *AppError {
	error := AppError{Message: message, StatusCode: statusCode, Details: details}
	return &error
}
