package router

import (
	c "blog-server-app/modules/comments/controllers"
	r "blog-server-app/modules/comments/repository"
	s "blog-server-app/modules/comments/services"
)

type CommentRouter struct {
	repository *r.CommentRepository
	service    *s.CommentService
	controller *c.CommentController

	*Router
}

func (router *CommentRouter) init() {

	router.repository = &r.CommentRepository{DB: router.DB, Logger: router.Logger.Named("CommentRepository")}

	router.service = &s.CommentService{CommentRepo: router.repository, Logger: router.Logger.Named("CommentService")}

	router.controller = &c.CommentController{CommentService: router.service, Logger: router.Logger.Named("CommentController")}

	router.mapRoute("/users/{userId}/blogs/{blogId}/comments/{id}", "GET", router.controller.GetCommentById)
	router.mapRoute("/users/{userId}/blogs/{blogId}/comments/{id}", "PATCH", router.controller.EditComment)
	router.mapRoute("/users/{userId}/blogs/{blogId}/comments/{id}", "DELETE", router.controller.DeleteComment)
	router.mapRoute("/users/{userId}/blogs/{blogId}/comments", "POST", router.controller.CreateComment)
}
