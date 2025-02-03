package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetUpMiddleware(app *fiber.App) {
	// CORS SetUp
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowOrigins: "http://localhost:3000", // Update this to match your frontend
	}))
}