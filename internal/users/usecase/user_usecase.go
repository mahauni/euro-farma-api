package usecase

import (
	"context"

	"github.com/mahauni/euro-farma-api/internal/users/entity"
)

type UserUsecase struct {
	repository entity.UserRepository
}

func NewCreateUserUseCase(repository entity.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (usecase *UserUsecase) CreateUser(ctx context.Context, inputUser *entity.User) error {
	user := entity.NewUser(inputUser.ID, inputUser.Name, inputUser.Password, inputUser.Email, inputUser.Points)

	err := usecase.repository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *UserUsecase) UpdateUser(ctx context.Context, inputUser *entity.User) error {
	err := usecase.repository.Update(ctx, inputUser)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	user, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.Update(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *UserUsecase) FindUserById(ctx context.Context, id int) (*entity.User, error) {
	user, err := usecase.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (usecase *UserUsecase) FindAllUsers(ctx context.Context) ([]entity.User, error) {
	users, err := usecase.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
