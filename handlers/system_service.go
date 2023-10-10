package handlers

import (
	"control-api/repositories"
	"github.com/gofiber/fiber/v2"
)

type SystemServiceHandler struct {
	app      *fiber.App
	userRepo *repositories.UserRepository
}

func NewSystemServiceHandler(app *fiber.App, userRepo *repositories.UserRepository) *SystemServiceHandler {
	serviceHandler := &SystemServiceHandler{
		app:      app,
		userRepo: userRepo,
	}
	app.Get("/services", serviceHandler.handleGetServices)
	app.Get("/services/:service", serviceHandler.handleGetService)
	app.Post("/services/:service/actions/:action", serviceHandler.handlePostServiceAction)
	return serviceHandler
}

func (handler *SystemServiceHandler) handleGetServices(ctx *fiber.Ctx) error {
	retrieve, err := handler.userRepo.Retrieve(2)
	if err != nil {
		return ctx.SendStatus(400)
	}
	return ctx.JSON(fiber.Map{"id": retrieve.Id})
}

func (handler *SystemServiceHandler) handleGetService(ctx *fiber.Ctx) error {
	serviceName := ctx.Params("service")
	return ctx.SendString(serviceName)
}

func (handler *SystemServiceHandler) handlePostServiceAction(ctx *fiber.Ctx) error {
	serviceName := ctx.Params("service")
	action := ctx.Params("action")
	return ctx.JSON(fiber.Map{"serviceName": serviceName, "action": action})
}
