package controllers

import (
	"blog-server-app/modules/blogs/models/dto"
	"blog-server-app/modules/blogs/services"
	errorHandler "blog-server-app/modules/system/handlers"
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
	response := controller.BlogService.CreateBlog(blog)
	return response, nil
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
