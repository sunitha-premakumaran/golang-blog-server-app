package controllers

import (
	"blog-server-app/internal/blogs/models/dto"
	"blog-server-app/internal/blogs/services"
	errorHandler "blog-server-app/pkg/handlers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//Controller is responsible for request validation and marshal and unmarshalling the response/request

type BlogController struct {
	BlogService *services.BlogService
	Logger      *zap.Logger
}

func (controller *BlogController) CreateBlog(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	var blog dto.CreateBlogDto
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Malformed Request", nil)
	}
	pathParams := mux.Vars(req)
	userId, ok := pathParams["userId"]
	if !ok {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Missing path param user id", nil)
	}
	return controller.BlogService.CreateBlog(userId, blog)
}

func (controller *BlogController) GetBlogById(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	pathParams := mux.Vars(req)
	id, ok := pathParams["id"]
	if !ok {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Missing path param blog id", nil)
	}
	return controller.BlogService.GetBlogById(id)
}

func (controller *BlogController) DeleteBlog(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	pathParams := mux.Vars(req)
	id, ok := pathParams["id"]
	if !ok {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Malformed Request", nil)
	}
	return controller.BlogService.DeleteBlogById(id)
}

func (controller *BlogController) EditBlog(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	pathParams := mux.Vars(req)
	id, ok := pathParams["id"]
	if !ok {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Missing path param blog id", nil)
	}
	var blog map[string]string
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Malformed Request", nil)
	}
	return controller.BlogService.UpdateBlogById(id, blog)
}
