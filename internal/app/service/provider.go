package service

import (
	"github.com/google/wire"
	"go-project-template/internal/app/service/foo"
)

var ProviderSet = wire.NewSet(
	NewBinding,
	foo.New,
)
