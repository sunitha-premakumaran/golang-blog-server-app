package controller

import (
	"blog-server-app/internal/comments/models/dto"
	"blog-server-app/internal/comments/services"
	errorHandler "blog-server-app/pkg/handlers"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CommentController struct {
	CommentService *services.CommentService
	Logger         *zap.Logger
}

func (controller *CommentController) CreateComment(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	var comment dto.CreateCommentDto
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Malformed Request", nil)
	}
	pathParams := mux.Vars(req)
	userId, userOk := pathParams["userId"]
	blogId, blogOk := pathParams["blogId"]
	if !userOk || !blogOk {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Missing path param user id or blog id", nil)
	}
	return controller.CommentService.CreateComment(userId, blogId, comment)
}

func (controller *CommentController) GetCommentById(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	pathParams := mux.Vars(req)
	blogId, blogOk := pathParams["blogId"]
	if !blogOk {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Missing path param user id or blog id", nil)
	}
	return controller.CommentService.GetComments(blogId)
}

func (controller *CommentController) DeleteComment(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	return nil, nil
}

func (controller *CommentController) EditComment(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	return nil, nil
}
