package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mahauni/fiap-gamify/internal/quizzes/entity"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type QuizRepositoryPostgres struct {
	db *pgx.Conn
	// db *pgxpool.Pool
}

func NewQuizRepositoryPostgres(db *pgx.Conn) *QuizRepositoryPostgres {
	return &QuizRepositoryPostgres{
		db: db,
	}
}

func (repo *QuizRepositoryPostgres) Create(ctx context.Context, quiz *entity.Quiz) error {
	query := `INSERT INTO quiz (question, answer) VALUES (@quizQuestion, @quizAnswer)`

	args := pgx.NamedArgs{
		"quizQuestion": quiz.Question,
		"quizAnswer":   quiz.Answers,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (repo *QuizRepositoryPostgres) FindById(ctx context.Context, id int) (*entity.Quiz, error) {
	query := `SELECT id, question, answer FROM quiz WHERE id = @quizId`

	args := pgx.NamedArgs{
		"quizId": id,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query quiz: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[*entity.Quiz])
}

func (repo *QuizRepositoryPostgres) FindAll(ctx context.Context) ([]*entity.Quiz, error) {
	query := `SELECT id, question, answer FROM quiz`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query quiz: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[*entity.Quiz])
}

func (repo *QuizRepositoryPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE * FROM quiz WHERE ID = @quizId`

	args := pgx.NamedArgs{
		"quizId": id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete quiz: %w", err)
	}

	return nil
}

func (repo *QuizRepositoryPostgres) Update(ctx context.Context, quiz *entity.Quiz) error {
	query := `UPDATE quiz SET question = @quizQuestion, answer = @quizAnswer, WHERE id = @quizId`

	args := pgx.NamedArgs{
		"quizQuestion": quiz.Question,
		"quizAnswer":   quiz.Answers,
		"quizId":       quiz.ID,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update quiz: %w", err)
	}

	return nil
}
