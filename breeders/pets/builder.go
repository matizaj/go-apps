package pets

import "errors"

type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(b string) *Pet
	SetMinWeight(w int) *Pet
	SetMaxWeight(w int) *Pet
	SetWeight(w int) *Pet
	SetDescription(d string) *Pet
	SetLifespan(l int) *Pet
	SetGeographicOrigin(g string) *Pet
	SetColor(c string) *Pet
	SetAge(a int) *Pet
	SetAgeEstimated(a bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}
func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}
func (p *Pet) SetBreed(breed string) *Pet {
	p.Breed = breed
	return p
}
func (p *Pet) SetColor(c string) *Pet {
	p.Color = c
	return p
}
func (p *Pet) SetDescription(d string) *Pet {
	p.Description = d
	return p
}
func (p *Pet) SetGeographicOrigin(g string) *Pet {
	p.GeographicOrigin = g
	return p
}
func (p *Pet) SetMinWeight(s int) *Pet {
	p.MinWeight = s
	return p
}
func (p *Pet) SetMaxWeight(s int) *Pet {
	p.MaxWeight = s
	return p
}
func (p *Pet) SetWeight(s int) *Pet {
	p.Weight = s
	return p
}

func (p *Pet) SetLifespan(s int) *Pet {
	p.Lifespan = s
	return p
}

func (p *Pet) SetAge(s int) *Pet {
	p.Age = s
	return p
}

func (p *Pet) SetAgeEstimated(s bool) *Pet {
	p.AgeEstimated = s
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("min weught cant be greather than max weight")
	}
	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}
