package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mahauni/euro-farma-api/internal/users/entity"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryPostgres struct {
	db *pgx.Conn
	// db *pgxpool.Pool
}

func NewUserRepositoryPostgres(db *pgx.Conn) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		db: db,
	}
}

func (repo *UserRepositoryPostgres) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (name, email, password, points) 
		VALUES (@userName, @userEmail, @userPassword, @userPoints)`

	args := pgx.NamedArgs{
		"userName":     user.Name,
		"userEmail":    user.Email,
		"userPassword": user.Password,
		"userPoints":   user.Points,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (repo *UserRepositoryPostgres) FindById(ctx context.Context, id int) (entity.User, error) {
	query := `SELECT id, name, email, password, points FROM users WHERE id = @userId`

	args := pgx.NamedArgs{
		"userId": id,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return entity.User{}, fmt.Errorf("unable to query user: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
}

func (repo *UserRepositoryPostgres) FindAll(ctx context.Context) ([]entity.User, error) {
	query := `SELECT id, name, email, password, points FROM users`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query user: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.User])
}

func (repo *UserRepositoryPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE * FROM users WHERE ID = @userId`

	args := pgx.NamedArgs{
		"userId": id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete user: %w", err)
	}

	return nil
}

func (repo *UserRepositoryPostgres) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET name = @userName, email = @userEmail, password = @userPassword, 
		points, @userPoints, WHERE id = @userId`

	args := pgx.NamedArgs{
		"userName":     user.Name,
		"userEmail":    user.Email,
		"userPassword": user.Password,
		"userPoints":   user.Points,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update user: %w", err)
	}

	return nil
}
