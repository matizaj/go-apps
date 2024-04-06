package user

import (
	"database/sql"
	"fmt"
	"github.com/matizaj/go-app/e-com/types"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("select * from users where email=?", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User) // empty pointer to user
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("select * from users where id=?", id)
	if err != nil {
		return nil, err
	}
	u := new(types.User) // empty pointer to user
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}
func (s *Store) CreateUser(user types.User) error {
	log.Println("create user")
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Println("err", err)
		return err
	}

	return nil
}
