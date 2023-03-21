package router

import (
	"blog-server-app/modules/blogs/models/dto"
	"blog-server-app/modules/system/handlers"

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

func (routeObj *Router) mapRoute(url string, restMethods string, f CustomHandlerFunc) {
	routeObj.Logger.Info(fmt.Sprintf("Mapped %s to %s", url, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()))
	routeObj.Router.HandleFunc(url, routeObj.requestHandler(f)).Methods(restMethods)
}

func (routeObj *Router) requestHandler(fn CustomHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := fn(w, r)

		statusCode := http.StatusOK
		correlationId := r.Header.Get("X-Correlation-ID")

		if err != nil {
			response, statusCode = routeObj.handleHTTPError(err, correlationId)
		} else {
			if r.Method == "POST" {
				statusCode = http.StatusCreated
			}
		}
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		routeObj.Logger.Info(fmt.Sprintf("[correlation-id:%s] %s<< %s %d %s", correlationId, r.Method, r.URL.Path, statusCode, http.StatusText(statusCode)))
	}

}

func (routeObj *Router) handleHTTPError(err error, correlationId string) (e *handlers.AppError, statusCode int) {
	switch e := err.(type) {
	case *handlers.AppError:
		e.CorrelationId = correlationId
		routeObj.Logger.Error(fmt.Sprintf("HTTP %d - %s", e.StatusCode, e.Message))
		return e, e.StatusCode
	default:
		routeObj.Logger.Error("HTTP unknown error occurred" + e.Error())
		return handlers.NewHTTPError(http.StatusInternalServerError, "Internal Server Error occured:"+e.Error(), nil), http.StatusInternalServerError
	}
}
