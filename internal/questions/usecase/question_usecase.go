package usecase

import (
	"context"

	"github.com/mahauni/euro-farma-api/internal/questions/entity"
)

type QuestionUsecase struct {
	repository entity.QuestionRepository
}

func NewCreateQuestionUseCase(repository entity.QuestionRepository) *QuestionUsecase {
	return &QuestionUsecase{
		repository: repository,
	}
}

func (usecase *QuestionUsecase) CreateQuestion(ctx context.Context, inputQuestion *entity.Question) error {
	question := entity.NewQuestion(inputQuestion.ID, inputQuestion.Question, inputQuestion.Score, inputQuestion.QuizId)

	err := usecase.repository.Create(ctx, question)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *QuestionUsecase) UpdateQuestion(ctx context.Context, inputQuestion *entity.Question) error {
	err := usecase.repository.Update(ctx, inputQuestion)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *QuestionUsecase) DeleteQuestion(ctx context.Context, id int) error {
	question, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, &question)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *QuestionUsecase) FindQuestionById(ctx context.Context, id int) (*entity.Question, error) {
	question, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (usecase *QuestionUsecase) FindAllQuestions(ctx context.Context) ([]entity.Question, error) {
	questions, err := usecase.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (usecase *QuestionUsecase) FindAllQuestionsByQuiz(ctx context.Context, quizId int) ([]entity.Question, error) {
	questions, err := usecase.repository.FindAllByQuiz(ctx, quizId)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
