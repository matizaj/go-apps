package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type LogsRepository interface {
	AddLogEntry(data any) error
}

type logRepository struct {
	DB *mongo.Client
}

func NewLogRepository(conn *mongo.Client) LogsRepository {
	return &logRepository{DB: conn}
}
