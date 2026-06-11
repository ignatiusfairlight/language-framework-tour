package posts

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	posts := GetAllPosts()
	c.JSON(http.StatusOK, posts)
}

func PostShow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		// Call PostShowService here
		"message": "I will show you one thing!",
	})
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