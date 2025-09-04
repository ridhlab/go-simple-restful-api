package answer

import "github.com/gofiber/fiber/v2"

type AnswerController struct {
	useCase *AnswerUseCase
}

func NewAnswerController(useCase *AnswerUseCase) *AnswerController {
	return &AnswerController{
		useCase: useCase,
	}
}

func (c *AnswerController) CreateAnswer(ctx *fiber.Ctx) error {
	return ctx.SendString("Create Answer")
}

func (c *AnswerController) GetAllAnswer(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Answer")
}

func (c *AnswerController) GetAnswerByID(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Answer By ID")
}

func (c *AnswerController) UpdateAnswer(ctx *fiber.Ctx) error {
	return ctx.SendString("Update Answer")
}

func (c *AnswerController) DeleteAnswer(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete Answer")
}
