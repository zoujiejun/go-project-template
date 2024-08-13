package application

import (
	"go-project-template/internal/pkg/httpServer"
)

type Runner struct {
	httpServer *httpServer.Server
}

func (r *Runner) getAppRunners() []AppRunner {
	return []AppRunner{
		r.httpServer,
	}
}

func NewRunner(httpServer *httpServer.Server) *Runner {
	return &Runner{httpServer: httpServer}
}
