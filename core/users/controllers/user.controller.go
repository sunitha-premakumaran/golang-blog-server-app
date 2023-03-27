package controller

import (
	"blog-server-app/internal/users/models/dto"
	"blog-server-app/internal/users/services"
	errorHandler "blog-server-app/pkg/handlers"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type UserController struct {
	UserService *services.UserService
	Logger      *zap.Logger
}

func (controller *UserController) CreateUser(resWriter http.ResponseWriter, req *http.Request) (interface{}, error) {
	var user dto.UserCreateDto
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		return nil, errorHandler.NewHTTPError(http.StatusBadRequest, "Malformed Request", nil)
	}
	return controller.UserService.CreateUser(user)
}
