package posts

import (
	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model

	Slug      string   `gorm:"uniqueIndex"`
	Title     string
	Content   string
}

func (PostModel) TableName() string {
	return "posts"
}
