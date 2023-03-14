package entities

import (
	"blog-server-app/modules/blogs/models/dto"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model

	Name        string
	Description string
	Tags        string
	Content     string
	Status      dto.BlogStatus
	CreatedBy   User      `gorm:"foreignKey:ID"`
	Comments    []Comment `gorm:"foreignKey:ID"`
}
