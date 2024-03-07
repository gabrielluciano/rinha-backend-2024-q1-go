package main

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/gabrielluciano/rinha-backend-2024-q1-go/config"
	"github.com/gabrielluciano/rinha-backend-2024-q1-go/internal/routes"
	"github.com/gabrielluciano/rinha-backend-2024-q1-go/internal/util"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	config.Connect()
	defer config.Pool.Close()

	app.Hooks().OnListen(func(data fiber.ListenData) error {
		go util.Warmup()
		return nil
	})

	app.Post("/clientes/:id/transacoes", routes.Transacao)
	app.Get("/clientes/:id/extrato", routes.Extrato)
	app.Get("/health", routes.Health)

	app.Listen(":" + os.Getenv("HTTP_PORT"))
}
