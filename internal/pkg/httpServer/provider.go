package httpServer

import (
	"github.com/google/wire"
	"go-project-template/internal/pkg/httpServer/middleware"
)

var ProviderSet = wire.NewSet(
	NewServer,
	NewRouter,
	middleware.ProviderSet,
)
