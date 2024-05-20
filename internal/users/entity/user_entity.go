package entity

import (
	"context"

	userTraining "github.com/mahauni/euro-farma-api/internal/userTraining/entity"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Points   int    `json:"points"`
	// i dont know if i need to do this
	Trainings []userTraining.UserTraining `json:"trainings"`
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
		Id:       id,
		Name:     name,
		Password: password,
		Email:    email,
		Points:   points,
	}
}
