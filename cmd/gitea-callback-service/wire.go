//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-project-template/internal/app/biz"
	"go-project-template/internal/app/data"
	"go-project-template/internal/app/service"
	"go-project-template/internal/pkg/application"
	"go-project-template/internal/pkg/config"
	"go-project-template/internal/pkg/database"
	"go-project-template/internal/pkg/httpServer"
	"go-project-template/internal/pkg/logger"
)

var providerSet = wire.NewSet(
	application.ProviderSet,
	config.ProviderSet,
	httpServer.ProviderSet,
	logger.ProviderSet,
	database.ProviderSet,

	service.ProviderSet,
	biz.ProviderSet,
	data.ProviderSet,
)

func InitializeApp(cfg string) (*application.Application, error) {
	panic(wire.Build(providerSet))
}
