package socrute

import (
	"github.com/gofiber/fiber/v2"
)

func emitir(c *fiber.Ctx) error {
	m := make(map[string]string)
	m["hola"] = "maquinas"

	return c.JSON(m)
}
