package entity

import "context"

type Answer struct {
	Id         int `json:"id"`
	Answer     string
	Correct    bool
	QuizId     int `json:"quiz_id" db:"quiz_id"`
	QuestionId int `json:"question_id" db:"question_id"`
}

type AnswerRepository interface {
	Create(ctx context.Context, quiz *Answer) error
	FindById(ctx context.Context, id int) (Answer, error)
	FindAll(ctx context.Context) ([]Answer, error)
	FindAllByQuiz(ctx context.Context, quizId int) ([]Answer, error)
	FindAllByQuestion(ctx context.Context, questionId int) ([]Answer, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Answer) error
}

func NewAnswer(id int, answer string, correct bool, quizId int, questionId int) *Answer {
	return &Answer{
		Id:         id,
		Answer:     answer,
		Correct:    correct,
		QuizId:     quizId,
		QuestionId: questionId,
	}
}
