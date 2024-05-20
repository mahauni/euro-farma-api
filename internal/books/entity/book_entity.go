package entity

import "context"

type Book struct {
	Id         int    `json:"id"`
	FilePath   string `json:"file_path"`
	FileName   string `json:"file_name"`
	TrainingId int    `json:"quiz_id" db:"quiz_id"`
}

type BookRepository interface {
	Create(ctx context.Context, quiz *Book) error
	FindById(ctx context.Context, id int) (Book, error)
	FindAll(ctx context.Context) ([]Book, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Book) error
}

func NewBook(id int, filePath string, fileName string, trainingId int) *Book {
	return &Book{
		Id:         id,
		FilePath:   filePath,
		FileName:   fileName,
		TrainingId: trainingId,
	}
}
