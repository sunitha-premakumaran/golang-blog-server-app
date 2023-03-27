package services

import (
	"blog-server-app/internal/users/models/dto"
	"blog-server-app/internal/users/repository"

	"go.uber.org/zap"
)

type UserService struct {
	UserRepository *repository.UserRepository
	Logger         *zap.Logger
}

func (service *UserService) CreateUser(user dto.UserCreateDto) (*dto.UserCreateResponseDto, error) {
	return service.UserRepository.CreateUser(user)
}
