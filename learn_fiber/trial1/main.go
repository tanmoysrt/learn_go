package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		Prefork: true,
		ReadBufferSize: 4096,
		WriteBufferSize: 4096,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:param", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	log.Fatal(app.Listen(":3000"))
}
