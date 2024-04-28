package entity

import "context"

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Points   int    `json:"points"`
}

type UserRepository interface {
	Create(ctx context.Context, quiz *User) error
	FindById(ctx context.Context, id int) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *User) error
}

func NewUser(id int, name string, password string, email string, points int) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: password,
		Email:    email,
		Points:   points,
	}
}
