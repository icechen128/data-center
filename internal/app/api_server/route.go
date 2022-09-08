package api_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icechen128/data-center/internal/app/api_server/handler"
)

func routeApi(route fiber.Router) {
	route.Get("/:workspace/:db/:table", handler.HandleList)
}
