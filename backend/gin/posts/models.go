package posts

import (
	"time"
)

type PostModel struct {
	ID         uint       `gorm:"primaryKey"`
	Title      string     `gorm:"column:title"` 
	Content    string     `gorm:"column:content"`
	CreatedAt  time.Time 
	UpdatedAt  time.Time
}

func (PostModel) TableName() string {
	return "posts"
}
