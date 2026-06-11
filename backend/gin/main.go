package main

import (
	"log"
	"net/http"

	"github.com/ignatiusfairlight/language-framework-tour/backend/gin/common"
	"github.com/ignatiusfairlight/language-framework-tour/backend/gin/posts"
	"github.com/gin-gonic/gin"
)

// func Migrate(db *gorm.DB) {}

func main() {

	db := common.InitDB()
	// Migrate(db)
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get to sqlDBQL:", err)
	} else {
		defer sqlDB.Close()
	}

	/*
	** Disabling automatic redirect for trailing slash
	*/
	r := gin.Default()
	r.RedirectTrailingSlash = false

	/*
	** Grouping the only routes available in this project
	*/
	v1 := r.Group("/api")
	posts.PostsRegister(v1.Group("/posts"))

	/*
	** Adding this test endpoint to ensure Gin itself works
	*/
	testPoint := r.Group("api/ping")

	testPoint.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8003")
}
