package models

type Answer struct {
	AnswerId   int    `json:"answer_id"`
	Content    string `json:"content"`
	AuthorId   int    `json:"author_id"`
	QuestionId int    `json:"question_id"`
	CreatedAt  string `json:"created_at"`
}
