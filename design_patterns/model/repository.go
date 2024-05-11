package model

import (
	"database/sql"
)

type CityRepository interface {
	GetAllCities() ([]*City, error)
}

type MysqlRepository struct {
	DB *sql.DB
}

func NewMysqlRepository(connection *sql.DB) CityRepository {
	return MysqlRepository{DB: connection}
}
