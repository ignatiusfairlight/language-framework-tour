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
	c.JSON(http.StatusCreated, gin.H{
		// Call PostCreateService here
		"message": "I will make you something!",
	})
}

func PostUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		// Call PostUpdateService here
		"message": "I will change something for you!",
	})
}

func PostDestroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		// Call PostDestroyService here
		"message": "I will destroy for you!",
	})
}