package posts

import "github.com/gin-gonic/gin"

type PostModelValidator struct {
	Title     string   `json:"title" binding:"required,max=255"`
	Content   string   `json:"content" binding:"required"`
}

func (v *PostModelValidator) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(v)
}