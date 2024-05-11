package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	maxOpenDbConn = 5
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

func InitDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("Failed to connect db: %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Panicf("Failed to ping db: %v", err)
		return nil, err
	}
	db.SetConnMaxLifetime(maxDbLifetime)
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)

	log.Println("Successfully connected to Db!")

	return db, nil
}
