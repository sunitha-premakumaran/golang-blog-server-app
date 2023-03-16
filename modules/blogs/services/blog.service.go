package services

import (
	"blog-server-app/modules/blogs/models/dto"
	"blog-server-app/modules/blogs/repository"
	errorHandler "blog-server-app/modules/system/handlers"

	"go.uber.org/zap"
)

// BlogService will have all the business logic for the CRUD operations

type BlogService struct {
	BlogRepo *repository.BlogRepo
	Logger   *zap.Logger
}

func (service *BlogService) CreateBlog(blog dto.CreateBlogDto) dto.CreateBlogResponseDto {
	return service.BlogRepo.CreateBlog(blog)
}

func (service *BlogService) GetBlogById(id string) (*dto.GetBlogDto, error) {
	blog := service.BlogRepo.GetBlogById(id)
	if blog.BlogId == 0 {
		return nil, errorHandler.NewHTTPError(404, "Blog not found", nil)
	}
	return &blog, nil
}

func (service *BlogService) DeleteBlogById(id string) (*dto.UpdateDeleteResponseDto, *errorHandler.AppError) {
	_, err := service.GetBlogById(id)

	if err != nil {
		return nil, errorHandler.NewHTTPError(404, "Blog not found", nil)
	}
	deleteResult := service.BlogRepo.DeleteBlog(id)
	return &deleteResult, nil
}

func (service *BlogService) UpdateBlogById(id string, patchProps map[string]string) (*dto.UpdateDeleteResponseDto, *errorHandler.AppError) {
	_, err := service.GetBlogById(id)

	if err != nil {
		return nil, errorHandler.NewHTTPError(404, "Blog not found", nil)
	}
	updateResult := service.BlogRepo.UpdateBlogById(id, patchProps)
	return &updateResult, nil
}
