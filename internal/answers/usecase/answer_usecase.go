package usecase

import (
	"context"

	"github.com/mahauni/euro-farma-api/internal/answers/entity"
)

type AnswerUsecase struct {
	repository entity.AnswerRepository
}

func NewCreateAnswerUseCase(repository entity.AnswerRepository) *AnswerUsecase {
	return &AnswerUsecase{
		repository: repository,
	}
}

func (usecase *AnswerUsecase) CreateAnswer(ctx context.Context, inputAnswer *entity.Answer) error {
	answer := entity.NewAnswer(
		inputAnswer.Id,
		inputAnswer.Answer,
		inputAnswer.Correct,
		inputAnswer.QuizId,
		inputAnswer.QuestionId,
	)

	err := usecase.repository.Create(ctx, answer)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *AnswerUsecase) UpdateAnswer(ctx context.Context, inputAnswer *entity.Answer) error {
	err := usecase.repository.Update(ctx, inputAnswer)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *AnswerUsecase) DeleteAnswer(ctx context.Context, id int) error {
	answer, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, &answer)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *AnswerUsecase) FindAnswerById(ctx context.Context, id int) (*entity.Answer, error) {
	answer, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &answer, nil
}

func (usecase *AnswerUsecase) FindAllAnswers(ctx context.Context) ([]entity.Answer, error) {
	answers, err := usecase.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return answers, nil
}

func (usecase *AnswerUsecase) FindAllAnswersByQuiz(ctx context.Context, quizId int) ([]entity.Answer, error) {
	answers, err := usecase.repository.FindAllByQuiz(ctx, quizId)
	if err != nil {
		return nil, err
	}

	return answers, nil
}

func (usecase *AnswerUsecase) FindAllAnswersByQuestion(ctx context.Context, questionId int) ([]entity.Answer, error) {
	answers, err := usecase.repository.FindAllByQuestion(ctx, questionId)
	if err != nil {
		return nil, err
	}

	return answers, nil
}
