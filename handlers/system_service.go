package handlers

import (
	"control-api/model"
	"control-api/repositories"
	"errors"
	"github.com/gofiber/fiber/v2"
	"os/exec"
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
	println(serviceName)
	println(action)
	rcon := new(model.Rcon)
	if err := ctx.BodyParser(rcon); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	println("command was: ", rcon.Command)
	out, err := exec.Command("cmd", "/c", rcon.Command).Output()

	if err != nil {
		exitCode := 0
		var exitError *exec.ExitError
		if errors.Is(err, exec.ErrDot) {

		}
		if errors.Is(err, exec.ErrNotFound) {

		}
		if errors.Is(err, exec.ErrWaitDelay) {

		}
		if errors.As(err, &exitError) {
			exitCode = exitError.ExitCode()
		}

		return ctx.Status(400).JSON(fiber.Map{"serviceName": serviceName, "action": action, "output": string(out), "exitCode": exitCode, "error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"serviceName": serviceName, "action": action, "output": string(out)})
}
