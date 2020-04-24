package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Fala Coders! Tudo certo?!")
	})

	app.Listen(3000)
}
