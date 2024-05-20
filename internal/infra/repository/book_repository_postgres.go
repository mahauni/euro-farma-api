package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mahauni/euro-farma-api/internal/books/entity"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type BookRepositoryPostgres struct {
	db *pgx.Conn
	// db *pgxpool.Pool
}

func NewBookRepositoryPostgres(db *pgx.Conn) *BookRepositoryPostgres {
	return &BookRepositoryPostgres{
		db: db,
	}
}

func (repo *BookRepositoryPostgres) Create(ctx context.Context, book *entity.Book) error {
	query := `INSERT INTO books (file_path, file_name, training_id) 
		VALUES (@bookFilePath, @bookFileName, @bookTrainingId)`

	args := pgx.NamedArgs{
		"bookFilePath":   book.FilePath,
		"bookFileName":   book.FileName,
		"bookTrainingId": book.TrainingId,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (repo *BookRepositoryPostgres) FindById(ctx context.Context, id int) (entity.Book, error) {
	query := `SELECT id, file_path, file_name, training_id FROM books WHERE id = @bookId`

	args := pgx.NamedArgs{
		"bookId": id,
	}

	rows, err := repo.db.Query(ctx, query, args)
	if err != nil {
		return entity.Book{}, fmt.Errorf("unable to query book: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Book])
}

func (repo *BookRepositoryPostgres) FindAll(ctx context.Context) ([]entity.Book, error) {
	query := `SELECT id, file_path, file_name, training_id FROM books`

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query book: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[entity.Book])
}

func (repo *BookRepositoryPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE * FROM books WHERE Id = @bookId`

	args := pgx.NamedArgs{
		"bookId": id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete book: %w", err)
	}

	return nil
}

func (repo *BookRepositoryPostgres) Update(ctx context.Context, book *entity.Book) error {
	query := `UPDATE books SET file_path = @bookFilePath, file_name = @bookFileName,
	training_id = @bookTrainingId WHERE id = @bookId`

	args := pgx.NamedArgs{
		"bookFilePath":   book.FilePath,
		"bookFileName":   book.FileName,
		"bookTrainingId": book.TrainingId,
		"bookId":         book.Id,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update book: %w", err)
	}

	return nil
}
