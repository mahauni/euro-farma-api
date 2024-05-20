package entity

import "context"

type Training struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
}

type TrainingRepository interface {
	Create(ctx context.Context, quiz *Training) error
	FindById(ctx context.Context, id int) (Training, error)
	FindAll(ctx context.Context) ([]Training, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Training) error
}

func NewTraining(id int, name string, categoryId int) *Training {
	return &Training{
		Id:         id,
		Name:       name,
		CategoryId: categoryId,
	}
}
