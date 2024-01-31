package router

import (
	"rest-api/domain/user/presentation"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(userController *presentation.UserController) *Router {
	app := fiber.New(fiber.Config{
		AppName: "Guja",
		Prefork: true,
	})

	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	v1 := app.Group("v1")
	v1.Post("/api/user", userController.Create)
	v1.Get("/api/user", userController.ReadAll)
	v1.Get("/api/user/:id", userController.Read)
	v1.Put("/api/user/:id", userController.Update)
	v1.Delete("/api/user/:id", userController.Delete)

	return &Router{App: app}
}

type Router struct {
	*fiber.App
}
