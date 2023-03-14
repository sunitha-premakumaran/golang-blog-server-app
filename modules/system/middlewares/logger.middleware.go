package middleware

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Logger is a middleware handler that does request logging
type LoggerMiddleware struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set correlation_id and print request
	correlation_id := r.Header.Get("X-Correlation-ID")
	if len(correlation_id) == 0 {
		correlation_id = uuid.New().String()
		r.Header.Set("X-Correlation-ID", correlation_id)
	}
	log.Printf("[correlation-id:%s] %s>> %s", correlation_id, r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

// NewLogger constructs a new Logger middleware handler
func NewLoggerMiddleware(handlerToWrap http.Handler) *LoggerMiddleware {
	return &LoggerMiddleware{handlerToWrap}
}
