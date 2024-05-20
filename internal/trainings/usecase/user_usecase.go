package usecase

import (
	"context"

	"github.com/mahauni/euro-farma-api/internal/trainings/entity"
)

type TrainingUsecase struct {
	repository entity.TrainingRepository
}

func NewCreateTrainingUseCase(repository entity.TrainingRepository) *TrainingUsecase {
	return &TrainingUsecase{
		repository: repository,
	}
}

func (usecase *TrainingUsecase) CreateTraining(ctx context.Context, inputTraining *entity.Training) error {
	training := entity.NewTraining(inputTraining.Id, inputTraining.Name, inputTraining.CategoryId)

	err := usecase.repository.Create(ctx, training)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrainingUsecase) UpdateTraining(ctx context.Context, inputTraining *entity.Training) error {
	err := usecase.repository.Update(ctx, inputTraining)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrainingUsecase) DeleteTraining(ctx context.Context, id int) error {
	training, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, &training)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrainingUsecase) FindTrainingById(ctx context.Context, id int) (*entity.Training, error) {
	training, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &training, nil
}

func (usecase *TrainingUsecase) FindAllTrainings(ctx context.Context) ([]entity.Training, error) {
	trainings, err := usecase.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return trainings, nil
}
