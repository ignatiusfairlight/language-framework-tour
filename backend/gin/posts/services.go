package posts

import (
    "github.com/ignatiusfairlight/language-framework-tour/backend/gin/common"
)

func GetAllPosts() []PostModel {
	db := common.GetDB()
	var posts []PostModel
	db.Find(&posts)
	return posts
}

func GetPostBySlug(condition interface{}) (PostModel, error) {
	db := common.GetDB()
	var posts PostModel
	err := db.Where(condition).First(&posts).Error
	return posts, err
}

func CreatePost(post *PostModel) error {
	db := common.GetDB()
	return db.Create(post).Error
} 

// func UpdatePost

// func DeletePost
