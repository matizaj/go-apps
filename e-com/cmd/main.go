package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/matizaj/go-app/e-com/cmd/api"
	"github.com/matizaj/go-app/e-com/config"
	data "github.com/matizaj/go-app/e-com/db"
	"log"
)

const webPort = "8099"

func main() {
	log.Println("start main fun")
	db, err := data.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	log.Println("b4 db")
	err = initStorage(db)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("afret db")

	server := api.NewApiServer(fmt.Sprintf(":%s", webPort), nil)
	err = server.Run()
	if err != nil {
		log.Panicln("Server cant start: ", err)
	}
}

func initStorage(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		log.Println("cant ping db", err)
		return err
	}
	log.Println("DB connected!")
	return nil
}
