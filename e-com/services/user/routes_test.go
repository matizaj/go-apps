package user

import (
	"bytes"
	"encoding/json"
	"github.com/matizaj/go-app/e-com/types"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockUserRepo struct{}

func (m *mockUserRepo) GetUserByEmail(email string) (*types.User, error) {
	u := types.User{
		FirstName: "mat",
		LastName:  "zaj",
		Email:     "asd@qwe.com",
		Password:  "pswd",
		CreatedAt: time.Now(),
	}
	return &u, nil
	//return nil, fmt.Errorf("user not found")
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
func (m *mockUserRepo) CreateUser(user types.User) error {
	return nil
}

func TestUserServiceHandler(t *testing.T) {
	userRepo := &mockUserRepo{}
	handler := NewHandler(userRepo)
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "test",
			LastName:  "test",
			Email:     "example",
			Password:  "password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected bad request code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}
