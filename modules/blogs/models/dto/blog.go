package dto

type BlogStatus string

const (
	Draft     BlogStatus = "draft"
	Published            = "published"
)

// CreateBlogDto
type CreateBlogDto struct {
	Name        string `json:"name" validate:"required"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	CreatedBy   User
	Status      BlogStatus `json:"status" validate:"required,oneof=draft published"`
}

type CreateBlogResponseDto struct {
	BlogId uint `json:"blogId"`
}

type GetBlogDto struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	CreatedBy   string `json:"createdBy"`
	BlogId      uint   `json:"blogId"`
	Status      string `json:"status"`
}

type UpdateDeleteResponseDto struct {
	AffectedRecords int64 `json:"affectedRecords"`
}
