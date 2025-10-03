package question

import (
	"errors"
	"github/com/ridhlab/go-simple-restful-api/internal/user"
	"github/com/ridhlab/go-simple-restful-api/models"
)

type IQuestionUseCase interface {
	CreateQuestion(question *models.Question) error
	GetQuestions() ([]*models.Question, error)
	GetQuestionById(id int) (*models.Question, error)
	GetQuestionByUserId(userId int) ([]*models.Question, error)
	UpdateQuestion(question *models.Question) error
	DeleteQuestion(id int) error
}

type QuestionUseCase struct {
	repo     IQuestionRepository
	userRepo user.IUserRepository
}

func NewQuestionUseCase(repo IQuestionRepository, userRepo user.IUserRepository) *QuestionUseCase {
	return &QuestionUseCase{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (uc *QuestionUseCase) CreateQuestion(createQuestion *CreateQuestionRequest) error {
	_, err := uc.userRepo.GetUserById(createQuestion.AuthorId)
	if err != nil {
		return errors.New("user not found")
	}
	err = uc.repo.CreateQuestion(&models.Question{
		Content:  createQuestion.Content,
		AuthorId: createQuestion.AuthorId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (uc *QuestionUseCase) GetQuestionById(id int) (*models.Question, error) {
	question, err := uc.repo.GetQuestionById(id)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (uc *QuestionUseCase) GetQuestionByUserId(userId int) ([]*models.Question, error) {
	questions, err := uc.repo.GetQuestionByUserId(userId)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (uc *QuestionUseCase) UpdateQuestion(updateQuestion *UpdateQuestionRequest) error {
	question, err := uc.repo.GetQuestionById(updateQuestion.QuestionId)
	if err != nil {
		return errors.New("question not found")
	}
	err = uc.repo.UpdateQuestion(&models.Question{
		QuestionId: updateQuestion.QuestionId,
		Content:    updateQuestion.Content,
		AuthorId:   question.AuthorId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (uc *QuestionUseCase) DeleteQuestion(id int) error {
	question, err := uc.repo.GetQuestionById(id)
	if err != nil {
		return errors.New("question not found")
	}
	err = uc.repo.DeleteQuestion(question.QuestionId)
	if err != nil {
		return err
	}
	return nil
}
