package foo

import (
	"github.com/gin-gonic/gin"
	"go-project-template/internal/app/biz/foo"
	"net/http"
	"strconv"
)

type Foo struct {
	biz *foo.Biz
}

func New(biz *foo.Biz) *Foo {
	return &Foo{biz: biz}
}

func (f *Foo) Foo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "bar")
}

func (f *Foo) List(ctx *gin.Context) {
	list, err := f.biz.GetFooList(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (f *Foo) Create(ctx *gin.Context) {
	var foo foo.Foo
	if err := ctx.ShouldBindJSON(&foo); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := f.biz.CreateFoo(ctx, &foo); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, foo)
}

func (f *Foo) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	foo, err := f.biz.GetFoo(ctx, idInt)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, foo)
}
