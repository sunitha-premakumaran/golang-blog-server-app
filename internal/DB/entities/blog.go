package entities

import (
	"blog-server-app/internal/blogs/models/dto"
	"database/sql"

	"github.com/lib/pq"
)

type Blog struct {
	BaseModel
	Name        string
	Description string
	Tags        pq.StringArray `gorm:"type:varchar(64)[]" json:"Tags"`
	Content     string
	Status      dto.BlogStatus
	AuthorID    sql.NullString
	Author      User `gorm:"foreignKey:AuthorID;references:ID"`
}
