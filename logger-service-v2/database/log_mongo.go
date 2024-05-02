package database

import (
	"context"
	"log"
)

func (l *logRepository) AddLogEntry(data any) error {
	collection := l.DB.Database("go-apps-logger").Collection("logs")

	insert, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("insertion id: ", insert.InsertedID)
	return nil
}
