package http

import (
	"github.com/gin-gonic/gin"
	"go-html-go/pkg/api"
)

type Router interface {
	routes() *gin.Engine
	Run(addr string) error
}

type router struct {
	api.Services
	routerEngine *gin.Engine
}

func NewRouter(r *gin.Engine, api api.Services) Router {
	return &router{api, r}
}

func (r *router) Run(addr string) error {
	routes := r.routes()

	return routes.Run(addr)
}
