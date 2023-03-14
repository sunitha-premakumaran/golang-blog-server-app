package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	content string
	Author  User `gorm:"foreignKey:ID"`
}
