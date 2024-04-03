package user

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

func (s *Store) getUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("select * from users where email=?", email)
	if err != nil {
		return nil, err
	}

}
