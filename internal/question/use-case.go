package question

import (
	"database/sql"
	"github/com/ridhlab/go-simple-restful-api/models"
)

type IQuestionUseCase interface {
	CreateQuestion(question *models.Question) error
	GetAllQuestion() ([]*models.Question, error)
	GetQuestionByID(id int) (*models.Question, error)
	UpdateQuestion(question *models.Question) error
	DeleteQuestion(id int) error
}

type QuestionUseCase struct {
	db *sql.DB
}

func NewQuestionUseCase(db *sql.DB) *QuestionUseCase {
	return &QuestionUseCase{
		db: db,
	}
}

func (uc *QuestionUseCase) CreateQuestion(question *models.Question) error {
	return nil
}

func (uc *QuestionUseCase) GetAllQuestion() ([]*models.Question, error) {
	return nil, nil
}

func (uc *QuestionUseCase) GetQuestionByID(id int) (*models.Question, error) {
	return nil, nil
}

func (uc *QuestionUseCase) UpdateQuestion(question *models.Question) error {
	return nil
}

func (uc *QuestionUseCase) DeleteQuestion(id int) error {
	return nil
}
