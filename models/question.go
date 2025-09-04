package models

type Question struct {
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
	AuthorId   int    `json:"author_id"`
	CreatedAt  string `json:"created_at"`
}
