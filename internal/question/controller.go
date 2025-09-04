package question

import "github.com/gofiber/fiber/v2"

type QuestionController struct {
	useCase *QuestionUseCase
}

func NewQuestionController(useCase *QuestionUseCase) *QuestionController {
	return &QuestionController{
		useCase: useCase,
	}
}

func (c *QuestionController) CreateQuestion(ctx *fiber.Ctx) error {
	return ctx.SendString("Create Question")
}

func (c *QuestionController) GetAllQuestion(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Question")
}

func (c *QuestionController) GetQuestionByID(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Question By ID")
}

func (c *QuestionController) UpdateQuestion(ctx *fiber.Ctx) error {
	return ctx.SendString("Update Question")
}

func (c *QuestionController) DeleteQuestion(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete Question")
}
