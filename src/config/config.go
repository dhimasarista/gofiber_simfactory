package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache"
)
s
// Setup initializes the Fiber app and applies configurations
func Setup(app *fiber.App) {
	// Set up Mustache template engine
	engine := mustache.New("./src/templates", ".mustache")
	app.Views(engine)

	// ... Other configurations ...

	// Start the server
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
