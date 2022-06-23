package plan

import (
	"github.com/gofiber/fiber/v2"
)

func lista(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "Lista"
	return c.JSON(m)
}
func user(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "user"
	return c.JSON(m)
}
func byname(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "byname"
	return c.JSON(m)
}
func byid(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "byid"
	return c.JSON(m)
}
func create(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "create"
	return c.JSON(m)
}
func delete(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "delete"
	return c.JSON(m)
}
func edit(c *fiber.Ctx) error {

	m := make(map[string]string)
	m["mensjae"] = "edit"
	return c.JSON(m)
}
