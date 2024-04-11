package types

import (
	"time"
)

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3"`
}
type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}
type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
}

type UserRepository interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user User) error
}
type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"qty"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductRepository interface {
	GetAllProducts() ([]Product, error)
	CreateProduct(product Product) error
	GetProductsByIds(ids []int) ([]Product, error)
}

type Order struct {
	Id        int       `json:"id"`
	UserId    int       `json:"uid"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	Id        int       `json:"id"`
	OrderId   int       `json:"oid"`
	ProductId int       `json:"pid"`
	Quantity  int       `json:"qty"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderRepository interface {
	CreateOrder(o Order) (int, error)
	CreateOrderItem(oi OrderItem) (int, error)
}
type CartItem struct {
	ProductId int
	Quantity  int
}
type CartCheckoutPayload struct {
	Items []CartItem `json:"cart_item" validate:"required"`
}
