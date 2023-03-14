package router

import (
	c "blog-server-app/modules/blogs/controllers"
	r "blog-server-app/modules/blogs/repository"
	s "blog-server-app/modules/blogs/services"
)

func (router *Router) initBlogRoutes() {

	repo := &r.BlogRepo{DB: router.DB}

	services := &s.BlogService{BlogRepo: repo}

	controller := c.Controller{BlogService: services}

	router.mapRoute("/users/{userId}/blogs/{id}", "GET", controller.GetBlogById)
	router.mapRoute("/users/{userId}/blogs/{id}", "PATCH", controller.EditBlog)
	router.mapRoute("/users/{userId}/blogs/{id}", "DELETE", controller.DeleteBlog)
	router.mapRoute("/users/{userId}/blogs", "POST", controller.CreateBlog)
}
