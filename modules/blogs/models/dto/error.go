package dto

type AppError struct {
	Message       string
	CorrelationId string
	StatusCode    int
	Details       string
}
