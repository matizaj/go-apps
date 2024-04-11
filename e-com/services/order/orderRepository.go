package order

import (
	"database/sql"
	"github.com/matizaj/go-app/e-com/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(o types.Order) (int, error) {
	query := `INSERT INTO orders (userId, total, status, address) VALUES (?, ?, ?, ?)`
	result, err := s.db.Exec(query, o.UserId, o.Total, o.Status, o.Address)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *Store) CreateOrderItem(oi types.OrderItem) (int, error) {
	query := `INSERT INTO order_items(orderId, productId, quantity, price) VALUES (?,?,?,?)`
	result, err := s.db.Exec(query, oi.OrderId, oi.ProductId, oi.Quantity, oi.Price)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
