package data

import (
	"github.com/google/wire"
	"go-project-template/internal/app/data/foo"
)

var ProviderSet = wire.NewSet(
	foo.New,
)
