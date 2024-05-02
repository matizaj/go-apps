package data

import (
	"github.com/matizaj/go-app/log-service-v2/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var repo database.LogsRepository

type Models struct {
	Log LogEntry
}
type LogEntry struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (l *LogEntry) AddLogEntry(data LogEntry) error {
	return repo.AddLogEntry(data)
}

func New(connection *mongo.Client) *Models {

	repo = database.NewLogRepository(connection)

	return &Models{
		Log: LogEntry{},
	}
}
