package entity

import "context"

type Quiz struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name"`
	TrainingId int    `json:"category"`
}

type QuizRepository interface {
	Create(ctx context.Context, quiz *Quiz) error
	FindById(ctx context.Context, id int) (Quiz, error)
	FindAll(ctx context.Context) ([]Quiz, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Quiz) error
}

func NewQuiz(id int, name string, trainingId int) *Quiz {
	return &Quiz{
		Id:         id,
		Name:       name,
		TrainingId: trainingId,
	}
}
