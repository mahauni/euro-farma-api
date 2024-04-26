package entity

import "context"

type Quiz struct {
	ID       int    `json:"id" db:"id"`
	Category string `json:"category"`
}

type QuizRepository interface {
	Create(ctx context.Context, quiz *Quiz) error
	FindById(ctx context.Context, id int) (Quiz, error)
	FindAll(ctx context.Context) ([]Quiz, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Quiz) error
}

func NewQuiz(id int, category string) *Quiz {
	return &Quiz{
		ID:       id,
		Category: category,
	}
}
