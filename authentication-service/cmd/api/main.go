package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/matizaj/go-app/authentication-service/data"
	"log"
	"net/http"
	"os"
	"time"
)

const webPort = "80"

var counts int64

type Config struct {
	Repo data.Repository
}

func main() {
	log.Println("Starting authentication service")

	//connect to DB
	//conn, err := connectDb()
	//if err != nil {
	//	log.Panic("cant connect db :(")
	//}

	//set app config
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDb(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectDb() (*sql.DB, error) {
	connectionString := os.Getenv("DSN")
	for {
		connection, err := openDb(connectionString)
		if err != nil {
			log.Println("postgres not yet ready...")
			counts++
		} else {
			log.Println("connected to postgres")
			return connection, nil
		}

		if counts > 10 {
			log.Println(err)
			return nil, err
		}
		log.Println("backing of for 2 seconds")
		time.Sleep(time.Second * 2)
	}
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostrgesRepo(conn)
	app.Repo = db
}
