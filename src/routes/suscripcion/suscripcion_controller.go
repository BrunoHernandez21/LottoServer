package suscripcion

import "github.com/gofiber/fiber/v2"

func crear(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
func eliminar(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
