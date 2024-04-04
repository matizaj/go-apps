package user

import (
	"github.com/matizaj/go-app/e-com/types"
	"testing"
	"time"
)

type mockUserRepo struct{}

func (m *mockUserRepo) GetUserByEmail(email string) (*types.User, error) {
	u := types.User{
		FirstName: "mat",
		LastName:  "zaj",
		Email:     "mat@zaj.com",
		Password:  "pswd",
		CreatedAt: time.Now(),
	}
	return &u, nil
}
func (m *mockUserRepo) GetUserById(id int) (*types.User, error) {
	u := types.User{
		FirstName: "mat",
		LastName:  "zaj",
		Email:     "mat@zaj.com",
		Password:  "pswd",
		CreatedAt: time.Now(),
	}
	return &u, nil
}
//func (m *mockUserRepo) CreaterUser/zzXx/x /xz (id int) (*types.User, error) {
//	u := types.User{
//		FirstName: "mat",
//		LastName:  "zaj",
		Email:     "mat@zaj.com",
		Password:  "pswd",
		CreatedAt: time.Now(),
	}
	return &u, nil
}

func TestUserServiceHandler(t *testing.T) {
	userRepo := &mockUserRepo
	handler := NewHandler(userRepo)
}
