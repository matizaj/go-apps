package data

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type Models struct {
	User User
}
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User: User{},
	}
}

func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at
				from users
				order by last_name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.FirstName,
			&u.LastName,
			&u.Password,
			&u.Active,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
