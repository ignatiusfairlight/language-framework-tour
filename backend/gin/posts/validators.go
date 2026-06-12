package posts

import "github.com/gin-gonic/gin"

type PostCreateValidator struct {
	Title     string   `json:"title" binding:"required,max=255"`
	Content   string   `json:"content" binding:"required"`
}

type PostUpdateValidator struct {
	Title     string   `json:"title" binding:"omitempty,max=255"`
	Content   string   `json:"content" binding:"omitempty"`
}

func (v *PostCreateValidator) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(v)
}

func (v *PostUpdateValidator) Bind (c *gin.Context) error {
	return c.ShouldBindJSON(v)
}