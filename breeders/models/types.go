package models

import (
	"database/sql"
	"time"
)

var repo Repository

type Models struct {
	DogBreed DogBreed
	CatBreed CatBreed
	Dog      Dog
}

func New(connection *sql.DB) *Models {
	if connection != nil {
		repo = newMysqlRepository(connection)
	} else {
		repo = newTestRepository(nil)
	}

	return &Models{
		DogBreed: DogBreed{},
	}
}
func NewTest() *Models {
	repo = newTestRepository(nil)

	return &Models{
		DogBreed: DogBreed{},
	}
}

type DogBreed struct {
	Id               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_heigh_lbs"`
	Lifespan         int    `json:"average_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type DogOfMonth struct {
	Id    int
	Dog   *Dog
	Video string
	Image string
}

func (d *DogBreed) GetAll() ([]*DogBreed, error) {
	return repo.AllDogBreeds()
}

func (d *DogBreed) GetBreedByName(b string) (*DogBreed, error) {
	return repo.GetBreedByName(b)
}

func (d *Dog) GetDogOfMonthById(id int) (*DogOfMonth, error) {
	return repo.GetDogOfMonthById(id)
}

//func (d *CatBreed) GetBreedByName(b string) (*CatBreed, error) {
//	return repo.GetBreedByName()
//}

type CatBreed struct {
	Id               int    `json:"id" xml:"id"`
	Breed            string `json:"breed" xml:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs" xml:"weight_low_lbs"`
	WeightMaxLbs     int    `json:"weight_max_lbs" xml:"weight_max_lbs"`
	AverageWeight    int    `json:"average_weight" xml:"average_weight"`
	Lifespan         int    `json:"average_lifespan" xml:"average_lifespan"`
	Details          string `json:"details" xml:"details"`
	AlternateNames   string `json:"alternate_names" xml:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin" xml:"geographic_origin"`
}
type Dog struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DayOfBirth       time.Time `json:"dob"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}
type Cat struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DayOfBirth       time.Time `json:"dob"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}
type Breeder struct {
	Id          int         `json:"id"`
	BreederName string      `json:"breeder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	Zip         string      `json:"zip"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}
type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"lifespan"`
}
