package models

func (m *testRepository) AllDogBreeds() ([]*DogBreed, error) {
	return nil, nil
}
func (m *testRepository) GetBreedByName(b string) (*DogBreed, error) {
	return &DogBreed{Breed: "Husky"}, nil
}
func (m *testRepository) GetDogOfMonthById(id int) (*DogOfMonth, error) { return nil, nil }
