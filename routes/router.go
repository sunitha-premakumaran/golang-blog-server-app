package router

import (
	"blog-server-app/modules/blogs/models/dto"
	"blog-server-app/modules/system/handlers"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"runtime"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Router struct {
	Router *mux.Router
	DB     *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	router := Router{Router: mux.NewRouter(), DB: db}

	// Healthz endpoint
	router.Router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		p := dto.HealthStatus{Status: "running", Health: "OK"}
		json.NewEncoder(w).Encode(p)
	})

	//Other modules
	router.initBlogRoutes()

	return &router
}

func (routeObj *Router) mapRoute(url string, restMethods string, f CustomHandlerFunc) {
	log.Printf("Mapped %s to %s", url, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	routeObj.Router.HandleFunc(url, makeErrorHandler(f)).Methods(restMethods)
}

type CustomHandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, interface{})

func makeErrorHandler(fn CustomHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := fn(w, r)
		if err != nil {

			switch e := err.(type) {
			case handlers.AppError:
				log.Printf("HTTP %d - %s", e.StatusCode, e.Message)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("X-Content-Type-Options", "nosniff")
				w.WriteHeader(e.StatusCode)
				json.NewEncoder(w).Encode(e)
			default:
				log.Println("HTTP unknown error occurred", e)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("X-Content-Type-Options", "nosniff")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(handlers.NewHTTPError(http.StatusInternalServerError, "Internal Server Error occured: Unknown", nil))
			}
		} else {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}
