package httpServer

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-project-template/internal/pkg/config"
	"go-project-template/internal/pkg/logger"
	"log"
	"net/http"
	"time"
)

type Server struct {
	serverConfig config.Server
	httpServer   *http.Server
	router       *Router

	logger *logger.Logger
}

func (s *Server) Name() string {
	return "http server"
}

func (s *Server) Start() error {
	if s.httpServer != nil {
		return errors.New("http server already started")
	}

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.serverConfig.Host, s.serverConfig.Port),
		Handler: s.router.engine,
	}

	go func() {
		s.logger.Info("Server started on port %d", s.serverConfig.Port)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %s", err.Error())
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	if s.httpServer == nil {
		return errors.New("http server not started")
	}

	s.logger.Info("http server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func NewServer(config *config.Config, router *Router, logger *logger.Logger) *Server {
	return &Server{
		serverConfig: config.Server,
		router:       router,
		logger:       logger,
	}
}
