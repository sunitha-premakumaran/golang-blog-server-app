package router

import (
	"blog-server-app/modules/blogs/models/dto"
	"blog-server-app/modules/system/handlers"
	appLogger "blog-server-app/modules/system/services"

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

func NewRouter(db *gorm.DB, logger *zap.Logger) *Router {
	router := Router{Router: mux.NewRouter(), DB: db, Logger: logger}

	// Healthz endpoint
	router.Router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		p := dto.HealthStatus{Status: "running", Health: "OK"}
		json.NewEncoder(w).Encode(p)
	})

	//Other modules
	router.initBlogRoutes()
	router.initCommentsRoutes()

	return &router
}

func (routeObj *Router) mapRoute(url string, restMethods string, f CustomHandlerFunc) {
	routeObj.Logger.Info(fmt.Sprintf("Mapped %s to %s", url, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()))
	routeObj.Router.HandleFunc(url, makeErrorHandler(f)).Methods(restMethods)
}

type CustomHandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func makeErrorHandler(fn CustomHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call controller
		response, err := fn(w, r)

		statusCode := http.StatusOK
		correlation_id := r.Header.Get("X-Correlation-ID")

		// Set common headers
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		logger := appLogger.NewAppLogger()

		if err != nil {
			switch e := err.(type) {
			case *handlers.AppError:
				statusCode = e.StatusCode
				logger.Error(fmt.Sprintf("HTTP %d - %s", e.StatusCode, e.Message))
				w.WriteHeader(e.StatusCode)
				e.CorrelationId = correlation_id
				json.NewEncoder(w).Encode(e)
			default:
				statusCode = http.StatusInternalServerError
				logger.Error("HTTP unknown error occurred" + e.Error())
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(handlers.NewHTTPError(http.StatusInternalServerError, "Internal Server Error occured: Unknown", nil))
			}
		} else {
			w.Header().Set("Content-type", "application/json")
			if r.Method == "POST" {
				statusCode = http.StatusCreated
			}
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(response)
		}
		logger.Info(fmt.Sprintf("[correlation-id:%s] %s<< %s %d %s", correlation_id, r.Method, r.URL.Path, statusCode, http.StatusText(statusCode)))
	}

}
