package question

import "github/com/ridhlab/go-simple-restful-api/models"

type CreateQuestionRequest struct {
	Content  string `json:"content"`
	AuthorId int    `json:"author_id"`
}

type UpdateQuestionRequest struct {
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
	AuthorId   int    `json:"author_id"`
}

type GetQuestionDetailResponse struct {
	models.Question
	Answers []*models.Answer `json:"answers"`
}
