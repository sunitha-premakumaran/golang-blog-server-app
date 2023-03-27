package repository

import (
	"blog-server-app/internal/DB/entities"
	"blog-server-app/internal/blogs/models/dto"
	"database/sql"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BlogRepo struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (repo *BlogRepo) CreateBlog(userId string, createDto dto.CreateBlogDto) (*dto.CreateBlogResponseDto, error) {
	response := dto.CreateBlogResponseDto{}
	dbDto := entities.Blog{Name: createDto.Name, Content: createDto.Content, Description: createDto.Description, Tags: createDto.Tags, Status: createDto.Status, AuthorID: sql.NullString{String: userId, Valid: true}}
	result := repo.DB.Create(&dbDto)
	if result.Error != nil && len(result.Error.Error()) != 0 {
		return nil, result.Error
	}
	response.BlogId = dbDto.ID
	return &response, nil
}

func (repo *BlogRepo) DeleteBlog(id string) (*dto.UpdateDeleteResponseDto, error) {
	blog := entities.Blog{}
	result := repo.DB.Delete(&blog, id)
	if result.Error != nil && len(result.Error.Error()) != 0 {
		return nil, result.Error
	}
	return &dto.UpdateDeleteResponseDto{AffectedRecords: result.RowsAffected}, nil
}

func (repo *BlogRepo) GetBlogById(id string) (*dto.GetBlogDto, error) {
	model := entities.Blog{}
	result := repo.DB.Where("ID = ?", id).First(&model)
	if result.Error != nil && len(result.Error.Error()) != 0 {
		return nil, result.Error
	}
	return &dto.GetBlogDto{Name: model.Name,
		Description: model.Description,
		Tags:        model.Tags,
		Content:     model.Content,
		Status:      string(model.Status),
		BlogId:      model.ID,
	}, nil
}

func (repo *BlogRepo) UpdateBlogById(id string, patchProps map[string]string) (*dto.UpdateDeleteResponseDto, error) {
	var keys []string
	for k := range patchProps {
		keys = append(keys, k)
	}
	blog := entities.Blog{}
	result := repo.DB.Model(&blog).Where("ID = ?", id).Update(keys[0], patchProps[keys[0]])
	if result.Error != nil && len(result.Error.Error()) != 0 {
		return nil, result.Error
	}
	return &dto.UpdateDeleteResponseDto{AffectedRecords: result.RowsAffected}, nil
}
