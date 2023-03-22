package repository

import (
	"blog-server-app/DB/entities"
	"blog-server-app/modules/comments/models/dto"
	"database/sql"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (repo *CommentRepository) CreateComment(userId string, blogId string, commentDto dto.CreateCommentDto) (*dto.CreateCommentResponseDto, error) {
	response := dto.CreateCommentResponseDto{}
	dbDto := entities.Comment{Content: commentDto.Content, PostedByUserID: sql.NullString{Valid: true, String: userId}, BelongsToBlogID: sql.NullString{Valid: true, String: blogId}}
	result := repo.DB.Create(&dbDto)
	if result.Error != nil && len(result.Error.Error()) != 0 {
		return nil, result.Error
	}
	response.CommentId = dbDto.ID
	return &response, nil
}

func (repo *CommentRepository) GetComments(blogId string) ([]entities.Comment, error) {
	comments := []entities.Comment{}
	result := repo.DB.First(&comments).Where("BlogId = ?", blogId)
	if result.Error != nil && len(result.Error.Error()) != 0 {
		return nil, result.Error
	}
	return comments, nil
}
