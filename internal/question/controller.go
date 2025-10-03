package question

import (
	"github/com/ridhlab/go-simple-restful-api/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type QuestionController struct {
	useCase *QuestionUseCase
}

func NewQuestionController(useCase *QuestionUseCase) *QuestionController {
	return &QuestionController{
		useCase: useCase,
	}
}

func (c *QuestionController) CreateQuestion(ctx *fiber.Ctx) error {
	var createQuestion CreateQuestionRequest
	if err := ctx.BodyParser(&createQuestion); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	err := c.useCase.CreateQuestion(&createQuestion)
	if err != nil {
		log.Printf("Error creating question: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create question",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Question created successfully",
	})
}

func (c *QuestionController) GetQuestions(ctx *fiber.Ctx) error {
	authorId := ctx.Query("author_id")
	var questions []*models.Question
	if authorId != "" {
		authorIdInt, err := strconv.Atoi(authorId)
		if err != nil {
			log.Printf("Error parsing author ID: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse author ID",
			})
		}

		questions, err = c.useCase.GetQuestionByUserId(authorIdInt)
		if err != nil {
			log.Printf("Error getting questions by user ID: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cannot get questions by user ID",
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"questions": questions,
		"message":   "Get question successfully",
	})
}

func (c *QuestionController) GetQuestionById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing question ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse question ID",
		})
	}
	question, err := c.useCase.GetQuestionById(id)
	if err != nil {
		log.Printf("Error getting question by ID: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get question by ID",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"question": question,
		"message":  "Get question by ID successfully",
	})
}

func (c *QuestionController) UpdateQuestion(ctx *fiber.Ctx) error {
	var updateQuestion UpdateQuestionRequest
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing question ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse question ID",
		})
	}
	updateQuestion.QuestionId = id
	if err := ctx.BodyParser(&updateQuestion); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	err = c.useCase.UpdateQuestion(&updateQuestion)
	if err != nil {
		log.Printf("Error updating question: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update question",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update question successfully",
	})
}

func (c *QuestionController) DeleteQuestion(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Printf("Error parsing question ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse question ID",
		})
	}
	err = c.useCase.DeleteQuestion(id)
	if err != nil {
		log.Printf("Error deleting question: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete question",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete question successfully",
	})
}
