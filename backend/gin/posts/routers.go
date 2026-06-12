package posts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostsRegister(router *gin.RouterGroup) {
	router.GET("/", PostIndex)
	router.GET("", PostIndex)
	router.GET("/:slug", PostShow)
	router.POST("/", PostCreate)
	router.POST("", PostCreate)
	router.PATCH("/:slug", PostUpdate)
	router.DELETE("/:slug", PostDestroy)
}

func PostIndex(c *gin.Context) {
	result := GetAllPosts()
	c.JSON(http.StatusOK, result)
}

func PostShow(c *gin.Context) {
	idStr := c.Param("slug")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := GetPostBySlug(&PostModel{ID: uint(id)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func PostCreate(c *gin.Context) {
	validator := PostCreateValidator{}
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := PostModel{
		Title: validator.Title,
		Content: validator.Content,
	}

	if err := CreatePost(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func PostUpdate(c *gin.Context) {
	idStr := c.Param("slug")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	validator := PostUpdateValidator{}
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := PostModel{
		Title: validator.Title,
		Content: validator.Content,
	}

	rows, err := UpdatePost(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post failed to be updated"})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated!"})
}

func PostDestroy(c *gin.Context) {
	idStr := c.Param("slug")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	rows, err := DeletePost(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Post could not be deleted"})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}