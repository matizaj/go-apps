package data

import (
	"database/sql"
	"time"
)

type PostgresTestRepo struct {
	Conn *sql.DB
}

func NewPostgresTestRepo(db *sql.DB) *PostgresTestRepo {
	return &PostgresTestRepo{Conn: db}
}

//func New(dbPool *sql.DB) Models {
//	db = dbPool
//	return Models{
//		User: User{},
//	}
//}

func (u *PostgresTestRepo) GetAll() ([]*User, error) {
	users := []*User{}
	return users, nil
}

func (u *PostgresTestRepo) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "First",
		LastName:  "Last",
		Password:  "password",
		Active:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return &user, nil
}

func (u *PostgresTestRepo) GetById(id int) (*User, error) {
	user := User{
		ID:        id,
		Email:     "test@example.com",
		FirstName: "First",
		LastName:  "Last",
		Password:  "password",
		Active:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return &user, nil
}

func (u *PostgresTestRepo) Update(user User) error {
	return nil
}

func (u *PostgresTestRepo) DeleteById(id int) error {
	return nil
}

func (u *PostgresTestRepo) Insert(user User) (int, error) {
	return 2, nil
}

func (u *PostgresTestRepo) ResetPassword(password string, user User) error {
	return nil
}

func (u *PostgresTestRepo) PasswordMatches(plainText string, user User) (bool, error) {
	return true, nil
}
