package biz

import (
	"github.com/google/wire"
	"go-project-template/internal/app/biz/foo"
)

var ProviderSet = wire.NewSet(
	foo.NewBiz,
)
