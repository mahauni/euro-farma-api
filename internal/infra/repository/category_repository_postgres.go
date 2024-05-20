package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mahauni/euro-farma-api/internal/categories/entity"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepositoryPostgres struct {
	db *pgx.Conn
	// db *pgxpool.Pool
}

func NewCategoryRepositoryPostgres(db *pgx.Conn) *CategoryRepositoryPostgres {
	return &CategoryRepositoryPostgres{
		db: db,
	}
}

func (repo *CategoryRepositoryPostgres) Create(ctx context.Context, category *entity.Category) error {
	query := `INSERT INTO categories (name) VALUES (@categoryName)`

	args := pgx.NamedArgs{
		"categoryName": category.Name,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (repo *CategoryRepositoryPostgres) FindById(ctx context.Context, id int) (entity.Category, error) {
	query := `SELECT id, name FROM categories WHERE id = @categoryId`

	args := pgx.NamedArgs{
		"categoryId": id,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return entity.Category{}, fmt.Errorf("unable to query category: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Category])
}

func (repo *CategoryRepositoryPostgres) FindAll(ctx context.Context) ([]entity.Category, error) {
	query := `SELECT id, name FROM categories`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query category: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Category])
}

func (repo *CategoryRepositoryPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE * FROM categories WHERE Id = @categoryId`

	args := pgx.NamedArgs{
		"categoryId": id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete category: %w", err)
	}

	return nil
}

func (repo *CategoryRepositoryPostgres) Update(ctx context.Context, category *entity.Category) error {
	query := `UPDATE categories SET name = @categoryName WHERE id = @categoryId`

	args := pgx.NamedArgs{
		"categoryName": category.Name,
		"categoryId":   category.Id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update category: %w", err)
	}

	return nil
}
