package httpServer

import (
	"github.com/gin-gonic/gin"
	"go-project-template/internal/pkg/config"
	"go-project-template/internal/pkg/httpServer/middleware"
)

type Router struct {
	engine *gin.Engine
	root   *gin.RouterGroup
}

func (r *Router) Root() *gin.RouterGroup {
	return r.root
}

type Binding interface {
	Bind(router *gin.RouterGroup)
}

func NewRouter(config *config.Config, binding Binding, middleware *middleware.GinMiddleware) *Router {
	gin.SetMode(config.Server.Mode)
	route := gin.New()
	route.Use(gin.Recovery())
	route.Use(middleware.RequestLogging.Middleware())

	r := &Router{
		engine: route,
		root:   route.Group("/"),
	}

	binding.Bind(r.root)

	return r
}
