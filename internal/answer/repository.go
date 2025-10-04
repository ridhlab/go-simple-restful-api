package answer

import (
	"database/sql"
	"github/com/ridhlab/go-simple-restful-api/models"
	"log"
)

type IAnswerRepository interface {
	CreateAnswer(answer *models.Answer) error
	GetAnswerById(id int) (*models.Answer, error)
	UpdateAnswer(answer *models.Answer) error
	DeleteAnswer(id int) error
	GetAnswerByQuestionId(id int) ([]*models.Answer, error)
	GetAnswerByUserId(id int) ([]*models.Answer, error)
}

type AnswerRepository struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) *AnswerRepository {
	return &AnswerRepository{
		db: db,
	}
}

func (r *AnswerRepository) CreateAnswer(answer *models.Answer) error {
	query := "INSERT INTO answers (question_id, author_id, content) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, answer.QuestionId, answer.AuthorId, answer.Content)
	if err != nil {
		log.Printf("Error creating answer: %v", err)
		return err
	}
	return nil
}

func (r *AnswerRepository) GetAnswerById(id int) (*models.Answer, error) {
	query := "SELECT id, question_id, author_id, content, created_at FROM answers WHERE id = $1"
	row := r.db.QueryRow(query, id)
	var answer models.Answer
	err := row.Scan(&answer.AnswerId, &answer.QuestionId, &answer.AuthorId, &answer.Content, &answer.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *AnswerRepository) GetAnswerByUserId(id int) ([]*models.Answer, error) {
	query := "SELECT id, question_id, author_id, content, created_at FROM answers WHERE author_id = $1"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var answers []*models.Answer
	for rows.Next() {
		var answer models.Answer
		err := rows.Scan(&answer.AnswerId, &answer.QuestionId, &answer.AuthorId, &answer.Content, &answer.CreatedAt)
		if err != nil {
			return nil, err
		}
		answers = append(answers, &answer)
	}
	return answers, nil
}

func (r *AnswerRepository) UpdateAnswer(answer *models.Answer) error {
	query := "UPDATE answers SET question_id = $1, author_id = $2, content = $3 WHERE id = $4"
	_, err := r.db.Exec(query, answer.QuestionId, answer.AuthorId, answer.Content, answer.AnswerId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AnswerRepository) DeleteAnswer(id int) error {
	query := "DELETE FROM answers WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AnswerRepository) GetAnswerByQuestionId(id int) ([]*models.Answer, error) {
	query := "SELECT id, question_id, author_id, content, created_at FROM answers WHERE question_id = $1"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var answers []*models.Answer
	for rows.Next() {
		var answer models.Answer
		err := rows.Scan(&answer.AnswerId, &answer.QuestionId, &answer.AuthorId, &answer.Content, &answer.CreatedAt)
		if err != nil {
			return nil, err
		}
		answers = append(answers, &answer)
	}
	return answers, nil
}
