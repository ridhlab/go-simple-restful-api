package user

import (
	"database/sql"
	"errors"
	"github/com/ridhlab/go-simple-restful-api/models"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	GetAllUser() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (fullname, email) VALUES ($1, $2) RETURNING user_id`
	_, err := r.db.Exec(query, user.Fullname, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAllUser() ([]*models.User, error) {
	query := `SELECT user_id, fullname, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.UserId, &user.Fullname, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserById(id int) (*models.User, error) {
	query := `SELECT user_id, fullname, email FROM users WHERE user_id = $1`
	row := r.db.QueryRow(query, id)
	user := &models.User{}
	if err := row.Scan(&user.UserId, &user.Fullname, &user.Email); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET fullname = $1, email = $2 WHERE user_id = $3`
	updatedData, err := r.db.Exec(query, user.Fullname, user.Email, user.UserId)
	if err != nil {
		return err
	}
	rowsAffected, err := updatedData.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE user_id = $1`
	updatedData, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := updatedData.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
