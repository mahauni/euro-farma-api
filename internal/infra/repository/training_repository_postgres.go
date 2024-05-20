package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mahauni/euro-farma-api/internal/trainings/entity"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type TrainingRepositoryPostgres struct {
	db *pgx.Conn
	// db *pgxpool.Pool
}

func NewTrainingRepositoryPostgres(db *pgx.Conn) *TrainingRepositoryPostgres {
	return &TrainingRepositoryPostgres{
		db: db,
	}
}

func (repo *TrainingRepositoryPostgres) Create(ctx context.Context, training *entity.Training) error {
	query := `INSERT INTO trainings (name, category_id) 
		VALUES (@trainingName, @trainingCategoryId)`

	args := pgx.NamedArgs{
		"trainingName":       training.Name,
		"trainingCategoryId": training.CategoryId,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (repo *TrainingRepositoryPostgres) FindById(ctx context.Context, id int) (entity.Training, error) {
	query := `SELECT id, name, category_id FROM trainings WHERE id = @trainingId`

	args := pgx.NamedArgs{
		"trainingId": id,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return entity.Training{}, fmt.Errorf("unable to query training: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Training])
}

func (repo *TrainingRepositoryPostgres) FindAll(ctx context.Context) ([]entity.Training, error) {
	query := `SELECT id, name, email, password, points FROM trainings`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query training: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Training])
}

func (repo *TrainingRepositoryPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE * FROM trainings WHERE Id = @trainingId`

	args := pgx.NamedArgs{
		"trainingId": id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete training: %w", err)
	}

	return nil
}

func (repo *TrainingRepositoryPostgres) Update(ctx context.Context, training *entity.Training) error {
	query := `UPDATE trainings SET name = @trainingName, category_id = @trainingCategoryId 
		WHERE id = @trainingId`

	args := pgx.NamedArgs{
		"trainingName":       training.Name,
		"trainingCategoryId": training.CategoryId,
		"trainingId":         training.Id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update training: %w", err)
	}

	return nil
}
