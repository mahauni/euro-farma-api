package entity

import "context"

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	Create(ctx context.Context, quiz *Category) error
	FindById(ctx context.Context, id int) (Category, error)
	FindAll(ctx context.Context) ([]Category, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, quiz *Category) error
}

func NewCategory(id int, name string) *Category {
	return &Category{
		Id:   id,
		Name: name,
	}
}
