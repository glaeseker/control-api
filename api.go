package main

import (
	"control-api/handlers"
	"control-api/model"
	"control-api/repositories"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
)

type APIServer struct {
	listenAddr string
	db         *sql.DB
}

func NewAPIServer(db *sql.DB, listenAddr string) *APIServer {
	return &APIServer{
		db:         db,
		listenAddr: listenAddr,
	}
}

func (api *APIServer) Run() {
	app := fiber.New()

	userRepo := repositories.NewUserRepository(api.db)

	handlers.NewSystemServiceHandler(app, userRepo)

	app.Get("/services/:service/journal", api.handleGetJournal)
	app.Get("/services/:service/journal/:page", api.handleGetJournalPage)
	app.Post("/services/:service/rcon", api.handlePostRcon)

	log.Println("Json API server running on port: ", api.listenAddr)
	app.Listen(api.listenAddr)
}

func (api *APIServer) handlePostRcon(c *fiber.Ctx) error {
	rcon := new(model.Rcon)
	if err := c.BodyParser(rcon); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.SendStatus(202)
}

func (api *APIServer) handleGetJournal(c *fiber.Ctx) error {
	return nil
}

func (api *APIServer) handleGetJournalPage(c *fiber.Ctx) error {
	return nil
}
