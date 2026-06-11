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

// func GetPostBySlug

// func CreatePost

// func UpdatePost

// func DeletePost
