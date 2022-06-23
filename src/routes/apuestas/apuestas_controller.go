package apuestas

import "github.com/gofiber/fiber/v2"

func crear(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
func activo(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
func lista(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
func activoPage(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
