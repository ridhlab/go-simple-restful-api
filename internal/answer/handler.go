package answer

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, controller *AnswerController) {
	answerGroup := app.Group("/answer")

	answerGroup.Get("/", func(c *fiber.Ctx) error {
		return controller.GetAnswers(c)
	})

	answerGroup.Get("/:id", func(c *fiber.Ctx) error {
		return controller.GetAnswerById(c)
	})

	answerGroup.Post("/", func(c *fiber.Ctx) error {
		return controller.CreateAnswer(c)
	})

	answerGroup.Put("/:id", func(c *fiber.Ctx) error {
		return controller.UpdateAnswer(c)
	})

	answerGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return controller.DeleteAnswer(c)
	})
}
