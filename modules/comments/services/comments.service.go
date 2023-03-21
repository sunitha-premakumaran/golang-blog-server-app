package services

import (
	"blog-server-app/DB/entities"
	"blog-server-app/modules/comments/models/dto"
	"blog-server-app/modules/comments/repository"

	"go.uber.org/zap"
)

type CommentService struct {
	CommentRepo *repository.CommentRepository
	Logger      *zap.Logger
}

func (service *CommentService) CreateComment(userId string, blogId string, createDto dto.CreateCommentDto) (*dto.CreateCommentResponseDto, error) {
	return service.CommentRepo.CreateComment(userId, blogId, createDto)
}

func (service *CommentService) GetComments(blogId string) ([]dto.GetCommentDto, error) {
	commentsDto, error := service.CommentRepo.GetComments(blogId)
	if len(error.Error()) != 0 {
		return nil, error
	}
	return transformDbCommentDto(commentsDto), nil
}

func transformDbCommentDto(comments []entities.Comment) []dto.GetCommentDto {
	commentsDto := []dto.GetCommentDto{}
	for _, comment := range comments {
		commentsDto = append(commentsDto, dto.GetCommentDto{Content: comment.Content})
	}
	return commentsDto
}
