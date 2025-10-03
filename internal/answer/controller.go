package answer

import (
	"github/com/ridhlab/go-simple-restful-api/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AnswerController struct {
	useCase *AnswerUseCase
}

func NewAnswerController(useCase *AnswerUseCase) *AnswerController {
	return &AnswerController{
		useCase: useCase,
	}
}

func (c *AnswerController) CreateAnswer(ctx *fiber.Ctx) error {
	var createAnswerReq CreateAnswerRequest
	err := ctx.BodyParser(&createAnswerReq)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.useCase.CreateAnswer(&createAnswerReq)
	if err != nil {
		log.Printf("Error creating answer: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.SendString("Create Answer")
}

func (c *AnswerController) GetAnswers(ctx *fiber.Ctx) error {
	questionId := ctx.Query("question_id")
	var answers []*models.Answer
	if questionId != "" {
		questionIdInt, err := strconv.Atoi(questionId)
		if err != nil {
			log.Printf("Error parsing question_id: %v", err)
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		answers, err = c.useCase.GetAnswerByQuestionId(questionIdInt)
		if err != nil {
			log.Printf("Error getting answers: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"answers": answers,
		"message": "Get answer successfully",
	})
}

func (c *AnswerController) GetAnswerById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	answer, err := c.useCase.GetAnswerById(id)
	if err != nil {
		log.Printf("Error getting answer: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"answer":  answer,
		"message": "Get answer successfully",
	})
}

func (c *AnswerController) UpdateAnswer(ctx *fiber.Ctx) error {
	var updateAnswerReq UpdateAnswerRequest
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updateAnswerReq.AnswerId = id
	err = ctx.BodyParser(&updateAnswerReq)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.useCase.UpdateAnswer(&updateAnswerReq)
	if err != nil {
		log.Printf("Error updating answer: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"answer":  updateAnswerReq,
		"message": "Update answer successfully",
	})
}

func (c *AnswerController) DeleteAnswer(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.useCase.DeleteAnswer(id)
	if err != nil {
		log.Printf("Error deleting answer: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete answer successfully",
	})
}
