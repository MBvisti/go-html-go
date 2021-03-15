package http

import (
	"github.com/gin-gonic/gin"
)

// routes holds all the routes you need for your service
func (r *router) routes() *gin.Engine {
	v1 := r.routerEngine

	v1.GET("/", RenderHome())

	return r.routerEngine
}
