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
	"time"
)

const webPort = "8088"

var counts int64

type Config struct {
	Repo   data.Repository
	Client *http.Client
}

func main() {
	log.Println("Starting authentication service")

	//connect to DB
	_, err := connectDb()
	if err != nil {
		log.Panic("cant connect db :(")
	}

	//set app config
	app := Config{
		Client: &http.Client{},
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
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
	log.Println("after ping db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectDb() (*sql.DB, error) {
	//connectionString := os.Getenv("DSN")
	connectionString := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
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
