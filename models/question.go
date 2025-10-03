package models

import "time"

type Question struct {
	QuestionId int       `json:"question_id"`
	Content    string    `json:"content"`
	AuthorId   int       `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
}
