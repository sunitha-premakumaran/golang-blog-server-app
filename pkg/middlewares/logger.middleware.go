package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Logger is a middleware handler that does request logging
type LoggerMiddleware struct {
	handler http.Handler
	Logger  *zap.Logger
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set correlation_id and print request
	correlationId := r.Header.Get("X-Correlation-ID")

	if correlationId == "" {
		// generate new version 4 uuid
		newid := uuid.New()
		correlationId = newid.String()
		r.Header.Set("X-Correlation-Id", correlationId)
		w.Header().Set("X-Correlation-Id", correlationId)
	}

	log.Printf("[correlation-id:%s] %s>> %s", correlationId, r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
	statusCode := 200
	log.Printf(fmt.Sprintf("[correlation-id:%s] %s<< %s %d %s", correlationId, r.Method, r.URL.Path, statusCode, http.StatusText(statusCode)))
}

// NewLogger constructs a new Logger middleware handler
func NewLoggerMiddleware(handlerToWrap http.Handler, logger *zap.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{handlerToWrap, logger}
}
