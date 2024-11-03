package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func ConnectionDB(cfg *mysql.Config) (*sql.DB, error) {
	connection, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("database connction error")
	}
	return connection, nil
}
