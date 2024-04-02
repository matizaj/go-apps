package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/matizaj/go-app/e-com/cmd/api"
	data "github.com/matizaj/go-app/e-com/db"
	"log"
)

const webPort = "8099"

func main() {
	db, err := data.NewMySqlStorage(mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 "127.0.1:3306",
		DBName:               "e-com",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NewApiServer(fmt.Sprintf(":%s", webPort), nil)
	err = server.Run()
	if err != nil {
		log.Panicln("Server cant start: ", err)
	}
}
