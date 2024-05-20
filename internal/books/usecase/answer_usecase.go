package usecase

import (
	"context"

	"github.com/mahauni/euro-farma-api/internal/books/entity"
)

type BookUsecase struct {
	repository entity.BookRepository
}

func NewCreateBookUseCase(repository entity.BookRepository) *BookUsecase {
	return &BookUsecase{
		repository: repository,
	}
}

func (usecase *BookUsecase) CreateBook(ctx context.Context, inputBook *entity.Book) error {
	book := entity.NewBook(
		inputBook.Id,
		inputBook.FilePath,
		inputBook.FileName,
		inputBook.TrainingId,
	)

	err := usecase.repository.Create(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *BookUsecase) UpdateBook(ctx context.Context, inputBook *entity.Book) error {
	err := usecase.repository.Update(ctx, inputBook)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *BookUsecase) DeleteBook(ctx context.Context, id int) error {
	book, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, &book)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *BookUsecase) FindBookById(ctx context.Context, id int) (*entity.Book, error) {
	book, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (usecase *BookUsecase) FindAllBooks(ctx context.Context) ([]entity.Book, error) {
	books, err := usecase.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
