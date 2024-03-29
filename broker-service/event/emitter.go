package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	log.Println("pushing to channel")

	err = channel.Publish("logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(event)},
	)
	if err != nil {
		return err
	}
	return nil
}

func NewEventEmitter(connection *amqp.Connection) (Emitter, error) {
	e := Emitter{
		connection: connection,
	}
	err := e.setup()
	if err != nil {
		return Emitter{}, err
	}
	return e, nil
}
