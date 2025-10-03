package user

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	useCase *UserUseCase
}

func NewUserController(useCase *UserUseCase) *UserController {
	return &UserController{
		useCase: useCase,
	}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var createUser CreateUserRequest
	if err := ctx.BodyParser(&createUser); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err := c.useCase.CreateUser(&createUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (c *UserController) GetAllUser(ctx *fiber.Ctx) error {
	users, err := c.useCase.GetAllUser()
	if err != nil {
		log.Printf("Error getting all user: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get all user",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get all user successfully",
		"users":   users,
	})
}

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	user, err := c.useCase.GetUserById(id)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get user by ID",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user by ID successfully",
		"user":    user,
	})
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	var updateUser UpdateUserRequest
	if err := ctx.BodyParser(&updateUser); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	updateUser.UserId = id
	err = c.useCase.UpdateUser(&updateUser)

	if err != nil {
		log.Printf("Error updating user: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update user",
		})
	}
	updatedUser, err := c.useCase.GetUserById(id)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get user by ID",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update user successfully",
		"user":    updatedUser,
	})
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	err = c.useCase.DeleteUser(id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete user",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete user successfully",
	})
}
