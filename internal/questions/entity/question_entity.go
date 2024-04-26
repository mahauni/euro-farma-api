package entity

import "context"

type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Score    int    `json:"score"`
	QuizId   int    `json:"quiz_id" db:"quiz_id"`
}

type QuestionRepository interface {
	Create(ctx context.Context, quiz *Question) error
	FindById(ctx context.Context, id int) (Question, error)
	FindAll(ctx context.Context) ([]Question, error)
	FindAllByQuiz(ctx context.Context, quizId int) ([]Question, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Question) error
}

func NewQuestion(id int, question string, score int, quizId int) *Question {
	return &Question{
		ID:       id,
		Question: question,
		Score:    score,
		QuizId:   quizId,
	}
}
