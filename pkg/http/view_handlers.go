package http

import (
	"github.com/gin-gonic/gin"
	view "go-html-go/pkg/views"
)

func RenderHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		view.Home(c.Writer)
	}
}
