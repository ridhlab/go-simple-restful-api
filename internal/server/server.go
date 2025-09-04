package server

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port string
	Host string
	App  *fiber.App
	db   *sql.DB
}

func NewServer(host, port string, db *sql.DB) *Server {
	app := fiber.New()
	server := Server{
		Port: port,
		Host: host,
		App:  app,
		db:   db,
	}

	return &server
}

func (s *Server) RunServer() error {
	log.Println("Starting server on " + s.Host + ":" + s.Port)
	s.RegisterRoutes()
	log.Fatal(s.App.Listen(s.Host + ":" + s.Port))
	return nil
}
