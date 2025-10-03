package user

import (
	"github/com/ridhlab/go-simple-restful-api/models"
)

type IUserUseCase interface {
	CreateUser(user *models.User) error
	GetAllUser() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type UserUseCase struct {
	repo IUserRepository
}

func NewUserUseCase(repo IUserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (uc *UserUseCase) CreateUser(cu *CreateUserRequest) error {
	err := uc.repo.CreateUser(&models.User{
		Fullname: cu.Fullname,
		Email:    cu.Email,
	})
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) GetAllUser() ([]*models.User, error) {
	users, err := uc.repo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUseCase) GetUserById(id int) (*models.User, error) {
	user, err := uc.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) UpdateUser(updatedUser *UpdateUserRequest) error {
	u, err := uc.repo.GetUserById(updatedUser.UserId)
	if err != nil {
		return err
	}

	if updatedUser.Fullname != nil {
		u.Fullname = *updatedUser.Fullname
	}
	if updatedUser.Email != nil {
		u.Email = *updatedUser.Email
	}

	err = uc.repo.UpdateUser(u)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) DeleteUser(id int) error {
	err := uc.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
