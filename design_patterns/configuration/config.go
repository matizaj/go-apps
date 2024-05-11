package configuration

import (
	"database/sql"
	"design-patterns/model"
	"sync"
)

type Application struct {
	Models *model.Model
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models: model.New(db),
		}
	})
	return instance
}
