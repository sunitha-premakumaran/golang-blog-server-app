package repository

import (
	"blog-server-app/DB/entities"
	"blog-server-app/modules/blogs/models/dto"

	"gorm.io/gorm"
)

type BlogRepo struct {
	DB *gorm.DB
}

func (repo *BlogRepo) CreateBlog(createDto dto.CreateBlogDto) dto.CreateBlogResponseDto {
	response := dto.CreateBlogResponseDto{}
	dbDto := entities.Blog{Name: createDto.Name, Content: createDto.Content, Description: createDto.Description, Tags: createDto.Tags, Status: createDto.Status}
	repo.DB.Create(&dbDto)
	response.BlogId = dbDto.ID
	return response
}

func (repo *BlogRepo) DeleteBlog(id string) dto.UpdateDeleteResponseDto {
	blog := entities.Blog{}
	result := repo.DB.Delete(&blog, id)
	return dto.UpdateDeleteResponseDto{AffectedRecords: result.RowsAffected}
}

func (repo *BlogRepo) GetBlogById(id string) dto.GetBlogDto {
	model := entities.Blog{}
	repo.DB.Where("ID = ?", id).First(&model)
	return dto.GetBlogDto{Name: model.Name,
		Description: model.Description,
		Tags:        model.Tags,
		Content:     model.Content,
		Status:      string(model.Status),
		BlogId:      model.ID,
	}
}

func (repo *BlogRepo) UpdateBlogById(id string, patchProps map[string]string) dto.UpdateDeleteResponseDto {
	var keys []string
	for k := range patchProps {
		keys = append(keys, k)
	}
	blog := entities.Blog{}
	result := repo.DB.Model(&blog).Where("ID = ?", id).Update(keys[0], patchProps[keys[0]])
	return dto.UpdateDeleteResponseDto{AffectedRecords: result.RowsAffected}
}

func GetBlogsByUser() {}
