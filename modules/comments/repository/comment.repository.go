package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (repo *CommentRepository) createComment() {

}
