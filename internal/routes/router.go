package router

import (
	"blog-server-app/internal/blogs/models/dto"
	"blog-server-app/pkg/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Router struct {
	Router *mux.Router
	DB     *gorm.DB
	Logger *zap.Logger
}

type BaseRouter interface {
	init()
}

func NewRouter(db *gorm.DB, logger *zap.Logger) *Router {
	router := Router{Router: mux.NewRouter(), DB: db, Logger: logger}

	// Healthz endpoint
	router.Router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		p := dto.HealthStatus{Status: "running", Health: "OK"}
		json.NewEncoder(w).Encode(p)
	})

	routers := []BaseRouter{&UserRouter{Router: &router}, &BlogRouter{Router: &router}, &CommentRouter{Router: &router}}

	for _, router := range routers {
		router.init()
	}

	return &router
}

type CustomHandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func (routeHandler *Router) mapRoute(url string, restMethods string, f CustomHandlerFunc) {
	routeHandler.Logger.Info(fmt.Sprintf("Mapped %s to %s", url, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()))
	routeHandler.Router.HandleFunc(url, routeHandler.requestHandler(f)).Methods(restMethods)
}

func (routeHandler *Router) requestHandler(fn CustomHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := fn(w, r)

		statusCode := http.StatusOK
		correlationId := r.Header.Get("X-Correlation-ID")

		if err != nil {
			response, statusCode = routeHandler.handleHTTPError(err, correlationId)
		} else {
			if r.Method == "POST" {
				statusCode = http.StatusCreated
			}
		}
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
	}
}

func (routeHandler *Router) handleHTTPError(err error, correlationId string) (e *handlers.AppError, statusCode int) {
	switch e := err.(type) {
	case *handlers.AppError:
		e.CorrelationId = correlationId
		routeHandler.Logger.Error(fmt.Sprintf("HTTP %d - %s", e.StatusCode, e.Message))
		return e, e.StatusCode
	default:
		routeHandler.Logger.Error("HTTP unknown error occurred" + e.Error())
		return handlers.NewHTTPError(http.StatusInternalServerError, "Internal Server Error occured:"+e.Error(), nil), http.StatusInternalServerError
	}
}
