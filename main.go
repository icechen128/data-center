package main

import (
	"github.com/icechen128/data-center/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(err)
		},
	})

	apiRoute := app.Group("/api")
	apiV1 := apiRoute.Group("/v1")

	apiV1.Get("/:workspace/:db/:table", api.HandleList)

	app.Listen(":8888")
}
