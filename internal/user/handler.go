package user

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, controller *UserController) {
	userGroup := app.Group("/user")

	userGroup.Get("/", func(c *fiber.Ctx) error {
		return controller.GetAllUser(c)
	})

	userGroup.Get("/:id", func(c *fiber.Ctx) error {
		return controller.GetUserByID(c)
	})

	userGroup.Post("/", func(c *fiber.Ctx) error {
		return controller.CreateUser(c)
	})

	userGroup.Put("/:id", func(c *fiber.Ctx) error {
		return controller.UpdateUser(c)
	})

	userGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return controller.DeleteUser(c)
	})
}
