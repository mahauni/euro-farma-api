package usecase

import (
	"context"

	"github.com/mahauni/fiap-gamify/internal/quizzes/entity"
)

type QuizUsecase struct {
	repository entity.QuizRepository
}

func NewCreateQuizUseCase(repository entity.QuizRepository) *QuizUsecase {
	return &QuizUsecase{
		repository: repository,
	}
}

func (usecase *QuizUsecase) CreateQuiz(ctx context.Context, inputQuiz *entity.Quiz) error {
	quiz := entity.NewQuiz(inputQuiz.ID, inputQuiz.Question, inputQuiz.Answers)

	err := usecase.repository.Create(ctx, quiz)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *QuizUsecase) UpdateQuiz(ctx context.Context, inputQuiz *entity.Quiz) error {
	quiz, err := usecase.repository.FindById(ctx, inputQuiz.ID)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, quiz)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *QuizUsecase) DeleteQuiz(ctx context.Context, id int) error {
	quiz, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, quiz)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *QuizUsecase) FindQuizById(ctx context.Context, id int) (*entity.Quiz, error) {
	quiz, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (usecase *QuizUsecase) FindAllQuiz(ctx context.Context) ([]*entity.Quiz, error) {
	quizzes, err := usecase.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return quizzes, nil
}
