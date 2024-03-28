package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

const webPort = "80"

type Config struct {
	Queue *amqp.Connection
}

func main() {

	rabbitConn, err := connect()
	if err != nil {
		log.Println("cant conect rabbitmq ", err)
		os.Exit(1)
	} else {
		log.Println("Connected to RabbitMQ")
	}
	defer rabbitConn.Close()

	log.Printf("Starting broker service on port %s\n", webPort)
	app := Config{
		Queue: rabbitConn,
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("cant strat broker sevice")
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = time.Second * 1
	var connection *amqp.Connection

	// dont continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("rabbitmq not ready")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("...backing off...")
		time.Sleep(backOff)

	}
	return connection, nil
}
