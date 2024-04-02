package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func NewMySqlStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	return db, nil
}
