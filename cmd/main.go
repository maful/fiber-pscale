package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/maful/fiber-pscale/handlers"
	"github.com/maful/fiber-pscale/models"
	"github.com/maful/fiber-pscale/util"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Fiber with Planetscale",
		ServerHeader: "Fiber",
	})

	if err := util.LoadConfig("."); err != nil {
		log.Fatal("cannot load config:", err)
	}

	models.ConnectDatabase()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "API with Fiber and PlanetScale",
		})
	})
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	app.Listen(":3000")
}
