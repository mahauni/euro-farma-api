package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mahauni/euro-farma-api/internal/questions/entity"
)

type QuestionRepositoryPostgres struct {
	db *pgx.Conn
	// db *pgxpool.Pool
}

func NewQuestionRepositoryPostgres(db *pgx.Conn) *QuestionRepositoryPostgres {
	return &QuestionRepositoryPostgres{
		db: db,
	}
}

func (repo *QuestionRepositoryPostgres) Create(ctx context.Context, question *entity.Question) error {
	query := `INSERT INTO quiz_questions (question, score, quiz_id) 
		VALUES (@questionQuestion, @questionScore, @questionQuizId)`

	args := pgx.NamedArgs{
		"questionQuestion": question.Question,
		"questionScore":    question.Score,
		"questionQuizId":   question.QuizId,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (repo *QuestionRepositoryPostgres) FindById(ctx context.Context, id int) (entity.Question, error) {
	query := `SELECT id, question, score, quiz_id FROM quiz_questions WHERE id = @questionId`

	args := pgx.NamedArgs{
		"questionId": id,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return entity.Question{}, fmt.Errorf("unable to query question: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Question])
}

func (repo *QuestionRepositoryPostgres) FindAll(ctx context.Context) ([]entity.Question, error) {
	query := `SELECT id, question, score, quiz_id FROM quiz_questions`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query question: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Question])
}

func (repo *QuestionRepositoryPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE * FROM quiz_questions WHERE Id = @questionId`

	args := pgx.NamedArgs{
		"questionId": id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete question: %w", err)
	}

	return nil
}

func (repo *QuestionRepositoryPostgres) Update(ctx context.Context, question *entity.Question) error {
	query := `UPDATE quiz_questions SET question = @questionQuestion, 
		score = @questionScore, quiz_id = @questionQuizId WHERE id = @questionId`

	args := pgx.NamedArgs{
		"questionQuestion": question.Question,
		"questionScore":    question.Score,
		"questionQuizId":   question.QuizId,
		"questionId":       question.Id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update question: %w", err)
	}

	return nil
}

func (repo *QuestionRepositoryPostgres) FindAllByQuiz(ctx context.Context, quizId int) ([]entity.Question, error) {
	query := `SELECT id, question, score, quiz_id FROM quiz_questions WHERE quiz_id = @questionQuizId`

	args := pgx.NamedArgs{
		"questionQuizId": quizId,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query question: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Question])
}
