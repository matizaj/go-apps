package models

import (
	"context"
	"log"
	"time"
)

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
				cast(((weight_low_lbs + weight_high_lbs) / 2)as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), 
				coalesce(geographic_origin, '')
				from dog_breeds order by breed`
	var breeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var d DogBreed
		err = rows.Scan(
			&d.Id,
			&d.Breed,
			&d.WeightLowLbs,
			&d.WeightHighLbs,
			&d.Lifespan,
			&d.Details,
			&d.AlternateNames,
			&d.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &d)
	}
	return breeds, nil
}

func (m *mysqlRepository) GetBreedByName(b string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select * from dog_breeds where breed=?`
	row := m.DB.QueryRowContext(ctx, query, b)

	var dogBreed DogBreed
	err := row.Scan(
		&dogBreed.Id,
		&dogBreed.Breed,
		&dogBreed.WeightLowLbs,
		&dogBreed.WeightHighLbs,
		&dogBreed.Lifespan,
		&dogBreed.Details,
		&dogBreed.AlternateNames,
		&dogBreed.GeographicOrigin,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &dogBreed, nil
}
func (m *mysqlRepository) GetDogOfMonthById(id int) (*DogOfMonth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	query := `select * from dog_of_month where id=?`
	row := m.DB.QueryRowContext(ctx, query, id)

	var dog DogOfMonth
	err := row.Scan(
		&dog.Id,
		&dog.Image,
		&dog.Video,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &dog, nil
}
