package products

import "time"

type Product struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) New() *Product {
	product := Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &product
}
