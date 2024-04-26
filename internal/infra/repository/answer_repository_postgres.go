package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mahauni/euro-farma-api/internal/answers/entity"
)

type AnswerRepositoryPostgres struct {
	db *pgx.Conn
	// db *pgxpool.Pool
}

func NewAnswerRepositoryPostgres(db *pgx.Conn) *AnswerRepositoryPostgres {
	return &AnswerRepositoryPostgres{
		db: db,
	}
}

func (repo *AnswerRepositoryPostgres) Create(ctx context.Context, answer *entity.Answer) error {
	query := `INSERT INTO quiz_answers (answer, correct, quiz_id, question_id) 
		VALUES (@answerAnswer, @answerCorrect, @answerQuizId, @answerQuestionId)`

	args := pgx.NamedArgs{
		"answerAnswer":     answer.Answer,
		"answerCorrect":    answer.Correct,
		"answerQuizId":     answer.QuizId,
		"answerQuestionId": answer.QuestionId,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (repo *AnswerRepositoryPostgres) FindById(ctx context.Context, id int) (entity.Answer, error) {
	query := `SELECT id, answer, correct, quiz_id, question_id FROM quiz_answers WHERE id = @answerId`

	args := pgx.NamedArgs{
		"answerId": id,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return entity.Answer{}, fmt.Errorf("unable to query answer: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Answer])
}

func (repo *AnswerRepositoryPostgres) FindAll(ctx context.Context) ([]entity.Answer, error) {
	query := `SELECT id, answer, correct, quiz_id, question_id FROM quiz_answers`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query answer: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Answer])
}

func (repo *AnswerRepositoryPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE * FROM quiz_answers WHERE ID = @answerId`

	args := pgx.NamedArgs{
		"answerId": id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete answer: %w", err)
	}

	return nil
}

func (repo *AnswerRepositoryPostgres) Update(ctx context.Context, answer *entity.Answer) error {
	query := `UPDATE quiz_answers SET answer = @answerAnswer, 
		correct = @answerCorrect, quiz_id = @answerQuizId, question_id = @answerQuestionId
		WHERE id = @answerId`

	args := pgx.NamedArgs{
		"answerAnswer":     answer.Answer,
		"answerCorrect":    answer.Correct,
		"answerQuizId":     answer.QuizId,
		"answerQuestionId": answer.QuestionId,
		"answerId":         answer.ID,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update answer: %w", err)
	}

	return nil
}

func (repo *AnswerRepositoryPostgres) FindAllByQuiz(ctx context.Context, quizId int) ([]entity.Answer, error) {
	query := `SELECT id, answer, correct, quiz_id, question_id 
		FROM quiz_answers WHERE quiz_id = @answerQuizId`

	args := pgx.NamedArgs{
		"answerQuizId": quizId,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query answer: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Answer])
}

func (repo *AnswerRepositoryPostgres) FindAllByQuestion(ctx context.Context, questionId int) ([]entity.Answer, error) {
	query := `SELECT id, answer, correct, quiz_id, question_id 
		FROM quiz_answers WHERE question_id = @answerQuestionId`

	args := pgx.NamedArgs{
		"answerQuestionId": questionId,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query answer: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Answer])
}
