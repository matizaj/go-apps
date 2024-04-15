package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	maxOpenDbCon  = 25
	maxIdleDbConn = 25
	maxDbLifetime = 5 * time.Minute
)

func initMySqlDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDbCon)
	db.SetConnMaxLifetime(maxDbLifetime)
	db.SetMaxIdleConns(maxIdleDbConn)

	return db, nil
}
