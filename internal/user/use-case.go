package user

import (
	"database/sql"
	"github/com/ridhlab/go-simple-restful-api/models"
)

type IUserUseCase interface {
	CreateUser(user *models.User) error
	GetAllUser() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type UserUseCase struct {
	db *sql.DB
}

func NewUserUseCase(db *sql.DB) *UserUseCase {
	return &UserUseCase{
		db: db,
	}
}

func (uc *UserUseCase) CreateUser(user *models.User) error {
	return nil
}

func (uc *UserUseCase) GetAllUser() ([]*models.User, error) {
	return nil, nil
}

func (uc *UserUseCase) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (uc *UserUseCase) UpdateUser(user *models.User) error {
	return nil
}

func (uc *UserUseCase) DeleteUser(id int) error {
	return nil
}
