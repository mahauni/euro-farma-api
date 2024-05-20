package entity

import "context"

type UserTraining struct {
	Id         int `json:"id" db:"id"`
	UserId     int `json:"user_id"`
	TrainingId int `json:"training_id"`
}

type UserTrainingRepository interface {
	Create(ctx context.Context, quiz *UserTraining) error
	FindById(ctx context.Context, id int) (UserTraining, error)
	FindAll(ctx context.Context) ([]UserTraining, error)
	FindAllByUser(ctx context.Context) ([]UserTraining, error)
	FindAllTraining(ctx context.Context) ([]UserTraining, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *UserTraining) error
}

func NewUserTraining(id int, userId int, trainingId int) *UserTraining {
	return &UserTraining{
		Id:         id,
		UserId:     userId,
		TrainingId: trainingId,
	}
}
