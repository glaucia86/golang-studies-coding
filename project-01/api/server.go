package main

import (
	"api-book/src"

	"github.com/gofiber/fiber"
)

//main ...função responsável por levantar servidor via Fiber
func main() {
	app := fiber.New()
	src.SetupRoutes(app)
	app.Listen("3000")
}
