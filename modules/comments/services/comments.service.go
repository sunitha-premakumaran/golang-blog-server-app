package services

import (
	"blog-server-app/modules/comments/repository"

	"go.uber.org/zap"
)

type CommentService struct {
	CommentRepo *repository.CommentRepository
	Logger      *zap.Logger
}
