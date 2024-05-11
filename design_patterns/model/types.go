package model

import (
	"database/sql"
	"log"
)

var repo CityRepository

type City struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	Area       int    `json:"ares"`
	Population int    `json:"population"`
}
type Model struct {
	City City
}

func New(c *sql.DB) *Model {
	repo = NewMysqlRepository(c)
	return &Model{
		City: City{},
	}
}
func (c *City) GetAll() ([]*City, error) {
	return repo.GetAllCities()
}

func (mysql MysqlRepository) GetAllCities() ([]*City, error) {
	query := `select * from cities`
	rows, err := mysql.DB.Query(query)
	if err != nil {
		log.Println("failed to query", err)
		return nil, err
	}
	defer rows.Close()

	var cities []*City
	for rows.Next() {
		var c City
		err = rows.Scan(
			&c.Name,
			&c.Country,
			&c.Area,
			&c.Population,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cities = append(cities, &c)
	}
	return cities, nil
}
