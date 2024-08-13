package service

import (
	"github.com/gin-gonic/gin"
	"go-project-template/internal/app/service/foo"
	"go-project-template/internal/pkg/httpServer"
)

type Binding struct {
	Foo *foo.Foo
}

func (b *Binding) Bind(router *gin.RouterGroup) {
	fooApi := router.Group("/foo")
	fooApi.GET("/", b.Foo.Foo)
	fooApi.POST("/", b.Foo.Create)
	fooApi.GET("/:id", b.Foo.Get)
	fooApi.GET("/sample/list", b.Foo.List)

}

func NewBinding(foo *foo.Foo) httpServer.Binding {
	return &Binding{Foo: foo}
}
