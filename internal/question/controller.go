package question

import (
	"github/com/ridhlab/go-simple-restful-api/internal/answer"
	"github/com/ridhlab/go-simple-restful-api/models"
	"log"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type QuestionController struct {
	useCase       *QuestionUseCase
	answerUseCase *answer.AnswerUseCase
}

func NewQuestionController(useCase *QuestionUseCase, answerUseCase *answer.AnswerUseCase) *QuestionController {
	return &QuestionController{
		useCase:       useCase,
		answerUseCase: answerUseCase,
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

	var question *models.Question
	var answers []*models.Answer

	var qErr, aErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		question, qErr = c.useCase.GetQuestionById(id)
	}()

	go func() {
		defer wg.Done()
		answers, aErr = c.answerUseCase.GetAnswerByQuestionId(id)
	}()

	wg.Wait()
	if qErr != nil {
		log.Printf("Error getting question by ID: %v", qErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get question by ID",
		})
	}
	if aErr != nil {
		log.Printf("Error getting answers by question ID: %v", aErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get answers by question ID",
		})
	}

	dataResult := &GetQuestionDetailResponse{
		Question: *question,
		Answers:  answers,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"question": dataResult,
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
