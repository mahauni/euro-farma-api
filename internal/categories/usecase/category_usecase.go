package usecase

import (
	"context"

	"github.com/mahauni/euro-farma-api/internal/categories/entity"
)

type CategoryUsecase struct {
	repository entity.CategoryRepository
}

func NewCreateCategoryUseCase(repository entity.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		repository: repository,
	}
}

func (usecase *CategoryUsecase) CreateCategory(ctx context.Context, inputCategory *entity.Category) error {
	category := entity.NewCategory(inputCategory.Id, inputCategory.Name)

	err := usecase.repository.Create(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *CategoryUsecase) UpdateCategory(ctx context.Context, inputCategory *entity.Category) error {
	err := usecase.repository.Update(ctx, inputCategory)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *CategoryUsecase) DeleteCategory(ctx context.Context, id int) error {
	category, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, &category)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *CategoryUsecase) FindCategoryById(ctx context.Context, id int) (*entity.Category, error) {
	category, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (usecase *CategoryUsecase) FindAllCategories(ctx context.Context) ([]entity.Category, error) {
	categories, err := usecase.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
