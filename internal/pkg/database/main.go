package database

import (
	"database/sql"
	"github.com/pkg/errors"
	"go-project-template/internal/pkg/config"
)

type Main struct {
	Driver string
	Conn   *sql.DB
}

func NewMain(config *config.Config) (*Main, error) {
	for _, cfg := range config.Databases {
		if cfg.Name == "main" {
			conn, err := connect(cfg)
			if err != nil {
				return nil, err
			}
			return &Main{
				Driver: cfg.Driver,
				Conn:   conn,
			}, nil
		}
	}

	return nil, errors.New("main database not found in config")
}
