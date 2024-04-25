package configuration

import (
	"database/sql"
	"github.com/matizaj/go-apps/go-breeders/adapters"
	"github.com/matizaj/go-apps/go-breeders/models"
	"sync"
)

type Application struct {
	Models     *models.Models
	CatService *adapters.RemoteService
}

var instance *Application
var once sync.Once
var db *sql.DB
var catService *adapters.RemoteService

func New(pool *sql.DB, cs *adapters.RemoteService) *Application {
	db = pool
	catService = cs
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models:     models.New(db),
			CatService: catService,
		}
	})
	return instance
}
