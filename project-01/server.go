package main

import (
	"github.com/gofiber/fiber"
)

//==> Função responsável por apresentar mensagem da API:
func mensagemAPIDefault(c *fiber.Ctx) {
	c.Send("Sejam bem-vindos(as) a API Golang + Fiber + PostGreSQL + Azure!")
}

//==> Função responsável por setar as rotas da aplicação:
func setupRoutes(app *fiber.App) {
	app.Get("/api", mensagemAPIDefault)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)
}
