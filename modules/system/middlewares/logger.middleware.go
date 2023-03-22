package middleware

import (
	"log"
	"net/http"
)

// Logger is a middleware handler that does request logging
type LoggerMiddleware struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set correlation_id and print request
	correlationId := r.Header.Get("X-Correlation-ID")
	log.Printf("[correlation-id:%s] %s>> %s", correlationId, r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

// NewLogger constructs a new Logger middleware handler
func NewLoggerMiddleware(handlerToWrap http.Handler) *LoggerMiddleware {
	return &LoggerMiddleware{handlerToWrap}
}
