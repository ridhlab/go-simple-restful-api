package question

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, controller *QuestionController) {
	questionGroup := app.Group("/question")

	questionGroup.Get("/", func(c *fiber.Ctx) error {
		return controller.GetQuestions(c)
	})

	questionGroup.Get("/:id", func(c *fiber.Ctx) error {
		return controller.GetQuestionById(c)
	})

	questionGroup.Post("/", func(c *fiber.Ctx) error {
		return controller.CreateQuestion(c)
	})

	questionGroup.Put("/:id", func(c *fiber.Ctx) error {
		return controller.UpdateQuestion(c)
	})

	questionGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return controller.DeleteQuestion(c)
	})
}
