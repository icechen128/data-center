package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleList(c *fiber.Ctx) error {
	workspace, err := MustParams(c, "workspace")
	if err != nil {
		return err
	}
	db, err := MustParams(c, "db")
	if err != nil {
		return err
	}
	table, err := MustParams(c, "table")
	if err != nil {
		return err
	}
	fmt.Println(strings.Join([]string{workspace, db, table}, ", "))

	return c.JSON(fiber.Map{"code": 0})
}

func MustParams(c *fiber.Ctx, name string) (value string, err error) {
	value = c.Params(name)
	if value == "" {
		err = errors.New("not found params of " + name)
	}
	return
}
