package router

import (
	c "blog-server-app/modules/blogs/controllers"
	r "blog-server-app/modules/blogs/repository"
	s "blog-server-app/modules/blogs/services"
)

func (router *Router) initBlogRoutes() {

	repo := &r.BlogRepo{DB: router.DB, Logger: router.Logger.Named("BlogRepository")}

	services := &s.BlogService{BlogRepo: repo, Logger: router.Logger.Named("BlogService")}

	controller := c.BlogController{BlogService: services, Logger: router.Logger.Named("BlogController")}

	router.mapRoute("/users/{userId}/blogs/{id}", "GET", controller.GetBlogById)
	router.mapRoute("/users/{userId}/blogs/{id}", "PATCH", controller.EditBlog)
	router.mapRoute("/users/{userId}/blogs/{id}", "DELETE", controller.DeleteBlog)
	router.mapRoute("/users/{userId}/blogs", "POST", controller.CreateBlog)
}
