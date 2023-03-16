package router

import (
	c "blog-server-app/modules/comments/controllers"
	r "blog-server-app/modules/comments/repository"
	s "blog-server-app/modules/comments/services"
)

func (router *Router) initCommentsRoutes() {

	repo := &r.CommentRepository{DB: router.DB, Logger: router.Logger.Named("CommentRepository")}

	services := &s.CommentService{CommentRepo: repo, Logger: router.Logger.Named("CommentService")}

	controller := c.CommentController{CommentService: services, Logger: router.Logger.Named("CommentController")}

	router.mapRoute("/users/{userId}/blogs/{blogId}/comments/{id}", "GET", controller.GetCommentById)
	router.mapRoute("/users/{userId}/blogs/{blogId}/comments/{id}", "PATCH", controller.EditComment)
	router.mapRoute("/users/{userId}/blogs/{blogId}/comments/{id}", "DELETE", controller.DeleteComment)
	router.mapRoute("/users/{userId}/blogs/{blogId}/comments", "POST", controller.CreateComment)
}
