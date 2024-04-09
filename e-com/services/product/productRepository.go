package product

import (
	"database/sql"
	"github.com/matizaj/go-app/e-com/types"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllProducts() ([]types.Product, error) {
	query := `select * from products`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var products []types.Product
	var p types.Product
	for rows.Next() {
		err = rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Image,
			&p.Price,
			&p.Quantity,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	products = append(products, p)
	return products, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	query := `INSERT INTO products(Name, Description, Image, Price, Quantity) VALUES (?,?,?,?,?)`
	result, err := s.db.Exec(query, product.Name, product.Description, product.Image, product.Price, product.Quantity)
	log.Println("result", result)
	if err != nil {
		log.Println("err", err)
		return err
	}

	return nil
}

func (s *Store) GetProductsByIds(ids []int) ([]types.Product, error) {
	query := `select * from products where id = ?`
	var p types.Product
	var pList []types.Product
	for _, id := range ids {
		rows, err := s.db.Query(query, id)
		for rows.Next() {
			err = rows.Scan(
				&p.Id,
				&p.Name,
				&p.Description,
				&p.Image,
				&p.Price,
				&p.Quantity,
				&p.CreatedAt,
			)
			if err != nil {
				return nil, err
			}
		}
		pList = append(pList, p)
	}
	return pList, nil
}
