package user

import "github.com/gofiber/fiber/v2"

type UserController struct {
	useCase *UserUseCase
}

func NewUserController(useCase *UserUseCase) *UserController {
	return &UserController{
		useCase: useCase,
	}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Create User")
}

func (c *UserController) GetAllUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All User")
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
	return ctx.SendString("Get User By ID")
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User")
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete User")
}
