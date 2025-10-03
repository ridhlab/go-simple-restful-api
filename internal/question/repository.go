package question

import (
	"database/sql"
	"fmt"
	"github/com/ridhlab/go-simple-restful-api/models"
)

type IQuestionRepository interface {
	CreateQuestion(question *models.Question) error

	GetQuestionById(id int) (*models.Question, error)
	GetQuestionByUserId(userId int) ([]*models.Question, error)
	UpdateQuestion(question *models.Question) error
	DeleteQuestion(id int) error
}

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) CreateQuestion(question *models.Question) error {
	query := `INSERT INTO questions (author_id, content) VALUES ($1, $2)`
	_, err := r.db.Exec(query, question.AuthorId, question.Content)
	if err != nil {
		return err
	}
	return nil
}

func (r *QuestionRepository) GetQuestionById(id int) (*models.Question, error) {
	query := `SELECT * FROM questions WHERE question_id = $1`
	row := r.db.QueryRow(query, id)
	var question models.Question
	if err := row.Scan(&question.QuestionId, &question.Content, &question.AuthorId, &question.CreatedAt); err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) GetQuestionByUserId(userId int) ([]*models.Question, error) {
	query := `SELECT * FROM questions WHERE author_id = $1`
	rows, err := r.db.Query(query, userId)
	fmt.Println(rows, err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*models.Question
	for rows.Next() {
		var question models.Question
		if err := rows.Scan(&question.QuestionId, &question.Content, &question.AuthorId, &question.CreatedAt); err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuestionRepository) UpdateQuestion(question *models.Question) error {
	query := `UPDATE questions SET content = $1, author_id = $2 WHERE question_id = $3`
	_, err := r.db.Exec(query, question.Content, question.AuthorId, question.QuestionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *QuestionRepository) DeleteQuestion(id int) error {
	query := `DELETE FROM questions WHERE question_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
