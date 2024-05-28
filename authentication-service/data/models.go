package data

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type PostgresRepo struct {
	Con *sql.DB
}

func NewPostrgesRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		Con: db,
	}
}

//	type Models struct {
//		User User
//	}
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

//func New(dbPool *sql.DB) Models {
//	db = dbPool
//	return Models{
//		User: User{},
//	}
//}

func (u *PostgresRepo) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at
				from public.users
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

func (u *PostgresRepo) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = $1`
	var user User
	row := db.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *PostgresRepo) GetById(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at 
				from users where id=$1`
	var user User
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *PostgresRepo) Update(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update users set
			emai; = $1,
			first_name=$2,
			last_name=$3,
			user_active=$4,
			updated_at=$5
			where id=$6
			`
	_, err := db.ExecContext(ctx, stmt, user.Email, user.FirstName, user.LastName, user.Active, time.Now(), user.ID)
	if err != nil {
		return err
	}
	return nil
}

//func (u *PostgresRepo) Delete() error {
//	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
//	defer cancel()
//
//	stmt := "delete from users where id=$1"
//	_, err := db.ExecContext(ctx, stmt, u.ID)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (u *PostgresRepo) DeleteById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := "delete from users where id=$1"
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *PostgresRepo) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int
	stmt := `insert into users(email, first_name, last_name, password, user_active, created_at, updated_at )
			values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err = db.QueryRowContext(ctx, stmt, user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now()).Scan(newID)
	if err != nil {
		return 0, nil
	}
	return newID, nil
}

func (u *PostgresRepo) ResetPassword(password string, user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `update users set password =$1 where id=$2`
	_, err = db.ExecContext(ctx, stmt, hashedPassword, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *PostgresRepo) PasswordMatches(plainText string, user User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, err
		default:
			return false, err
		}
	}
	return true, nil

}
