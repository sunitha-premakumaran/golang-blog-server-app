package controller

import (
	errorHandler "blog-server-app/modules/system/handlers"
	"blog-server-app/modules/users/models/dto"
	"blog-server-app/modules/users/services"
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
