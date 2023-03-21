package router

import (
	c "blog-server-app/modules/blogs/controllers"
	r "blog-server-app/modules/blogs/repository"
	s "blog-server-app/modules/blogs/services"
)

type BlogRouter struct {
	repository *r.BlogRepo
	service    *s.BlogService
	controller *c.BlogController

	*Router
}

func (router *BlogRouter) init() {

	router.repository = &r.BlogRepo{DB: router.DB, Logger: router.Logger.Named("BlogRepository")}

	router.service = &s.BlogService{BlogRepo: router.repository, Logger: router.Logger.Named("BlogService")}

	router.controller = &c.BlogController{BlogService: router.service, Logger: router.Logger.Named("BlogController")}

	router.mapRoute("/users/{userId}/blogs/{id}", "GET", router.controller.GetBlogById)
	router.mapRoute("/users/{userId}/blogs/{id}", "PATCH", router.controller.EditBlog)
	router.mapRoute("/users/{userId}/blogs/{id}", "DELETE", router.controller.DeleteBlog)
	router.mapRoute("/users/{userId}/blogs", "POST", router.controller.CreateBlog)
}
