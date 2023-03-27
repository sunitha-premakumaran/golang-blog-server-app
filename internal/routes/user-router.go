package router

import (
	c "blog-server-app/internal/users/controllers"
	r "blog-server-app/internal/users/repository"
	s "blog-server-app/internal/users/services"
)

type UserRouter struct {
	repository *r.UserRepository
	service    *s.UserService
	controller c.UserController

	*Router
}

func (router *UserRouter) init() {

	router.repository = &r.UserRepository{DB: router.DB, Logger: router.Logger.Named("UserRepository")}

	router.service = &s.UserService{UserRepository: router.repository, Logger: router.Logger.Named("UserService")}

	router.controller = c.UserController{UserService: router.service, Logger: router.Logger.Named("UserController")}

	router.mapRoute("/users", "POST", router.controller.CreateUser)
}
