package api_server

import (
	"github.com/gofiber/fiber/v2"
)

func RunServer() error {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(err)
		},
	})

	apiRoute := app.Group("/api")
	apiV1 := apiRoute.Group("/v1")
	routeApi(apiV1)

	return app.Listen(":8888")
}
