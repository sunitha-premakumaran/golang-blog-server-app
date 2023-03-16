package controller

import (
	"blog-server-app/modules/comments/services"
	"net/http"

	"go.uber.org/zap"
)

type CommentController struct {
	CommentService *services.CommentService
	Logger         *zap.Logger
}

func (controller *CommentController) CreateComment(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	return nil, nil
}

func (controller *CommentController) GetCommentById(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	return nil, nil

}

func (controller *CommentController) DeleteComment(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	return nil, nil
}

func (controller *CommentController) EditComment(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	return nil, nil
}
