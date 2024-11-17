package http

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func createServer(host string, port int) error {
	app := fiber.New(&fiber.Settings{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"status":  "success",
			"message": "hello from medomeckz!",
		})
	})

	address := fmt.Sprintf("%s:%d", host, port)
	return app.Listen(address)

}

func NewCreateServer(host string, port int) error {
	return createServer(host, port)
}
