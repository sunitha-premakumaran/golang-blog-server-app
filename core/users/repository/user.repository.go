package repository

import (
	"blog-server-app/internal/DB/entities"
	"blog-server-app/internal/users/models/dto"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (repo *UserRepository) CreateUser(user dto.UserCreateDto) (*dto.UserCreateResponseDto, error) {
	response := dto.UserCreateResponseDto{}
	dbDto := entities.User{}
	result := repo.DB.Create(&dbDto)
	if result.Error != nil && len(result.Error.Error()) != 0 {
		return nil, result.Error
	}
	response.CreatedUserId = dbDto.ID
	return &response, nil
}
