package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"nem"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func New(mongo *mongo.Client) Models {
	return Models{
		LogEntry: LogEntry{},
	}
}

func (l *LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logs_db").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: entry.CreatedAt,
		UpdatedAt: entry.UpdatedAt,
	})

	if err != nil {
		log.Println("error inserting into logs failed: ", err)
		return err
	}
	return nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	collection := client.Database("logs_db").Collection("logs")
	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs err: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var logs []*LogEntry

	for cursor.Next(ctx) {
		var item LogEntry
		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Failed to decode log item into slice: ", err)
			return nil, err
		} else {
			logs = append(logs, &item)
		}
	}
	return logs, nil
}
