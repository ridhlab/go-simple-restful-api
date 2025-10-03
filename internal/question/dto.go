package question

type CreateQuestionRequest struct {
	Content  string `json:"content"`
	AuthorId int    `json:"author_id"`
}

type UpdateQuestionRequest struct {
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
	AuthorId   int    `json:"author_id"`
}
