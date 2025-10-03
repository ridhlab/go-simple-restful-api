package answer

type CreateAnswerRequest struct {
	QuestionId int    `json:"question_id"`
	AuthorId   int    `json:"author_id"`
	Content    string `json:"content"`
}

type UpdateAnswerRequest struct {
	AnswerId int    `json:"answer_id"`
	Content  string `json:"content"`
}
