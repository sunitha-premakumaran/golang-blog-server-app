package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

// Logger is a middleware handler that does request logging
type CommonHeadersMiddleware struct {
	handler http.Handler
	Logger  *zap.Logger
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *CommonHeadersMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.handler.ServeHTTP(w, r)
	w.Header().Set("Content-type", "application/json")
}

// NewLogger constructs a new Logger middleware handler
func NewCommonHeadersMiddleware(handlerToWrap http.Handler, logger *zap.Logger) *CommonHeadersMiddleware {
	return &CommonHeadersMiddleware{handlerToWrap, logger}
}
