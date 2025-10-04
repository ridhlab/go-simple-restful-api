package answer

import (
	"errors"
	"github/com/ridhlab/go-simple-restful-api/internal/user"
	"github/com/ridhlab/go-simple-restful-api/models"
	"log"
)

type IAnswerUseCase interface {
	CreateAnswer(answer *models.Answer) error
	GetAnswerById(id int) (*models.Answer, error)
	GetAnswerByQuestionId(id int) ([]*models.Answer, error)
	GetAnswerByUserId(id int) ([]*models.Answer, error)
	UpdateAnswer(answer *models.Answer) error
	DeleteAnswer(id int) error
}

type AnswerUseCase struct {
	repo     IAnswerRepository
	userRepo user.IUserRepository
}

func NewAnswerUseCase(repo IAnswerRepository, userRepo user.IUserRepository) *AnswerUseCase {
	return &AnswerUseCase{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (uc *AnswerUseCase) CreateAnswer(createAnswerReq *CreateAnswerRequest) error {
	_, err := uc.userRepo.GetUserById(createAnswerReq.AuthorId)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return errors.New("user not found")
	}
	answer := &models.Answer{
		QuestionId: createAnswerReq.QuestionId,
		AuthorId:   createAnswerReq.AuthorId,
		Content:    createAnswerReq.Content,
	}
	err = uc.repo.CreateAnswer(answer)
	if err != nil {
		log.Printf("Error creating answer: %v", err)
		return err
	}
	return nil
}

func (uc *AnswerUseCase) GetAnswerById(id int) (*models.Answer, error) {
	answer, err := uc.repo.GetAnswerById(id)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (uc *AnswerUseCase) GetAnswerByQuestionId(id int) ([]*models.Answer, error) {
	answers, err := uc.repo.GetAnswerByQuestionId(id)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (uc *AnswerUseCase) GetAnswerByUserId(id int) ([]*models.Answer, error) {
	answers, err := uc.repo.GetAnswerByUserId(id)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (uc *AnswerUseCase) UpdateAnswer(updateAnswerReq *UpdateAnswerRequest) error {
	answer, err := uc.repo.GetAnswerById(updateAnswerReq.AnswerId)
	if err != nil {
		return err
	}
	answer.Content = updateAnswerReq.Content
	err = uc.repo.UpdateAnswer(answer)
	if err != nil {
		return err
	}
	return nil
}

func (uc *AnswerUseCase) DeleteAnswer(id int) error {
	_, err := uc.repo.GetAnswerById(id)
	if err != nil {
		return err
	}
	err = uc.repo.DeleteAnswer(id)
	if err != nil {
		return err
	}
	return nil
}
