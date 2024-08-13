package application

import (
	"fmt"
	"go-project-template/internal/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

type AppRunner interface {
	Start() error
	Stop() error
	Name() string
}

type Application struct {
	appName    string
	logger     *logger.Logger
	appRunners []AppRunner
}

func New(appName string, logger *logger.Logger, runner *Runner) *Application {
	return &Application{
		appName:    appName,
		logger:     logger,
		appRunners: runner.getAppRunners(),
	}
}

func (app *Application) Start() error {
	for _, runner := range app.appRunners {
		if err := runner.Start(); err != nil {
			return err
		} else {
			app.logger.Info(fmt.Sprintf("application %s started", runner.Name()))
		}
	}
	return nil
}

func (app *Application) AwaitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	app.logger.Info("application is shutting down")
	for _, runner := range app.appRunners {
		if err := runner.Stop(); err != nil {
			app.logger.Error(fmt.Sprintf("application %s stopped with error: %s", runner.Name(), err.Error()))
		}
	}
	os.Exit(0)
}
