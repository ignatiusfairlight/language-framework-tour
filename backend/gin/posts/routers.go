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
	c.JSON(http.StatusOK, gin.H{
		"message": "I will show you everything!",
	})
}

func PostShow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I will show you one thing!",
	})
}

func PostCreate(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "I will make you something!",
	})
}

func PostUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I will change something for you!",
	})
}

func PostDestroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I will destroy for you!",
	})
}