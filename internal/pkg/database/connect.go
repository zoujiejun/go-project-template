package database

import (
	"database/sql"
	"fmt"
	"github.com/didi/gendry/manager"
	"github.com/didi/gendry/scanner"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"go-project-template/internal/pkg/config"
	"sync"
)

var once sync.Once

func connect(config config.Database) (db *sql.DB, err error) {
	defer func() {
		if err == nil {
			db.SetMaxIdleConns(config.MaxIdleConns)
			db.SetMaxOpenConns(config.MaxOpenConns)
		}
	}()

	once.Do(func() {
		scanner.SetTagName("structs")
	})

	switch config.Driver {
	case "mysql":
		if config.Dsn != "" {
			return sql.Open(config.Driver, config.Dsn)
		}

		settings := make([]manager.Setting, 0)
		if config.Location == "" {
			config.Location = "Asia/Shanghai"
		}
		settings = append(settings, manager.SetLoc(config.Location))
		if config.Charset == "" {
			config.Charset = "utf8mb4"
		}
		settings = append(settings, manager.SetCharset(config.Charset))
		settings = append(settings, manager.SetParseTime(config.ParseTime))

		return manager.New(config.Database, config.Username, config.Password, config.Host).Set(settings...).Port(config.Port).Open(true)
	case "sqlite3":
		return sql.Open(config.Driver, config.Dsn)
	}

	return nil, errors.New(fmt.Sprintf("database driver %s is not supported", config.Driver))
}
