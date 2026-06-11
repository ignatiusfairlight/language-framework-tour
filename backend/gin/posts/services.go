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

func UpdatePost(id uint, data *PostModel) (int64, error) {
	db := common.GetDB()
	result := db.Model(&PostModel{ID: id}).Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
 
func DeletePost(id uint) (int64, error) {
	db := common.GetDB()
	result := db.Delete(&PostModel{}, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
