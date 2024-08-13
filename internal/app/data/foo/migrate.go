package foo

import (
	"database/sql"
	"github.com/pkg/errors"
)

const (
	mysqlDDL = `CREATE TABLE IF NOT EXISTS foo (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	)`
	sqliteDDL = `CREATE TABLE IF NOT EXISTS foo (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL
	)`
)

func migrate(driverName string, db *sql.DB) error {
	var err error
	switch driverName {
	case "mysql":
		_, err = db.Exec(mysqlDDL)
	case "sqlite3":
		_, err = db.Exec(sqliteDDL)
	default:
		return errors.New("unsupported driver")
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
