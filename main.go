package main

import (
	"data-center/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	apiRoute := app.Group("/api")
	apiV1 := apiRoute.Group("/v1")

	apiV1.Get("/:workspace/:db/:table", api.HandleList)

	app.Listen(":8888")
}
