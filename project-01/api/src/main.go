package src

import (
	"github.com/gofiber/fiber"
)

// SetupRoutes função responsável por setar as rotas da aplicação:
func SetupRoutes(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Seja bem-vindos(as) a API - Golang + Fiber + PostGreSQL + Azure!")
	})
}
