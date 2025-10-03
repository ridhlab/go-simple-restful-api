package server

import (
	"fmt"
	"github/com/ridhlab/go-simple-restful-api/internal/answer"
	"github/com/ridhlab/go-simple-restful-api/internal/question"
	"github/com/ridhlab/go-simple-restful-api/internal/user"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) RegisterRoutes() {
	fmt.Println("Registering routes...")
	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Welcome to the Go Simple RESTful API!")
	})

	userRepo := user.NewUserRepository(s.db)
	userUseCase := user.NewUserUseCase(userRepo)
	userController := user.NewUserController(userUseCase)
	user.RegisterRoutes(s.App, userController)

	answerRepo := answer.NewAnswerRepository(s.db)
	answerUseCase := answer.NewAnswerUseCase(answerRepo, userRepo)
	answerController := answer.NewAnswerController(answerUseCase)
	answer.RegisterRoutes(s.App, answerController)

	questionRepo := question.NewQuestionRepository(s.db)
	questionUseCase := question.NewQuestionUseCase(questionRepo, userRepo)
	questionController := question.NewQuestionController(questionUseCase)
	question.RegisterRoutes(s.App, questionController)
}
