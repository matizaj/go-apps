package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// connect rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println("cant conect rabbitmq ", err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ")
	// start listening for messages
	// create consumer - consume massages from queue
	// watch queue - consume events

}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = time.Second * 1
	var connection *amqp.Connection

	// dont continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
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
