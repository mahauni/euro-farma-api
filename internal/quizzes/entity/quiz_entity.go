package entity

import "context"

type Answers struct {
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}

type Quiz struct {
	ID       int       `json:"id" db:"id"`
	Question string    `json:"question"`
	Answers  []Answers `json:"answers"`
}

type QuizRepository interface {
	Create(ctx context.Context, quiz *Quiz) error
	FindById(ctx context.Context, id int) (*Quiz, error)
	FindAll(ctx context.Context) ([]*Quiz, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Quiz) error
}

func NewQuiz(id int, question string, answers []Answers) *Quiz {
	return &Quiz{
		ID:       id,
		Question: question,
		Answers:  answers,
	}
}
