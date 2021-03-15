package http

import (
	"go-html-go/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// add static dir
	routes.StaticFS("/static", http.Dir("./pkg/views/static"))

	return routes.Run(addr)
}
