package entities

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
