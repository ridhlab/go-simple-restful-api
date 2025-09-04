package answer

import (
	"database/sql"
	"github/com/ridhlab/go-simple-restful-api/models"
)

type IAnswerUseCase interface {
	CreateAnswer(answer *models.Answer) error
	GetAllAnswer() ([]*models.Answer, error)
	GetAnswerByID(id int) (*models.Answer, error)
	UpdateAnswer(answer *models.Answer) error
	DeleteAnswer(id int) error
}

type AnswerUseCase struct {
	db *sql.DB
}

func NewAnswerUseCase(db *sql.DB) *AnswerUseCase {
	return &AnswerUseCase{
		db: db,
	}
}

func (uc *AnswerUseCase) CreateAnswer(answer *models.Answer) error {
	return nil
}

func (uc *AnswerUseCase) GetAllAnswer() ([]*models.Answer, error) {
	return nil, nil
}

func (uc *AnswerUseCase) GetAnswerByID(id int) (*models.Answer, error) {
	return nil, nil
}

func (uc *AnswerUseCase) UpdateAnswer(answer *models.Answer) error {
	return nil
}

func (uc *AnswerUseCase) DeleteAnswer(id int) error {
	return nil
}
