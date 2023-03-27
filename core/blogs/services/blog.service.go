package services

import (
	"blog-server-app/internal/blogs/models/dto"
	"blog-server-app/internal/blogs/repository"
	errorHandler "blog-server-app/pkg/handlers"

	"go.uber.org/zap"
)

// BlogService will have all the business logic for the CRUD operations

type BlogService struct {
	BlogRepo *repository.BlogRepo
	Logger   *zap.Logger
}

func (service *BlogService) CreateBlog(userId string, blog dto.CreateBlogDto) (*dto.CreateBlogResponseDto, error) {
	return service.BlogRepo.CreateBlog(userId, blog)
}

func (service *BlogService) GetBlogById(id string) (*dto.GetBlogDto, error) {
	blog, err := service.BlogRepo.GetBlogById(id)
	if err != nil {
		return nil, err
	}
	if len(blog.BlogId) == 0 {
		return nil, errorHandler.NewHTTPError(404, "Blog not found", nil)
	}
	return blog, nil
}

func (service *BlogService) DeleteBlogById(id string) (*dto.UpdateDeleteResponseDto, error) {
	_, err := service.GetBlogById(id)

	if err != nil {
		return nil, errorHandler.NewHTTPError(404, "Blog not found", nil)
	}
	return service.BlogRepo.DeleteBlog(id)
}

func (service *BlogService) UpdateBlogById(id string, patchProps map[string]string) (*dto.UpdateDeleteResponseDto, error) {
	_, err := service.GetBlogById(id)

	if err != nil {
		return nil, errorHandler.NewHTTPError(404, "Blog not found", nil)
	}
	return service.BlogRepo.UpdateBlogById(id, patchProps)
}
