package suscripcion

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func listar(c *fiber.Ctx) error {
	input := []gormdb.Suscripciones{}
	db.Find(&input, "Usuario_id = ?", c.Locals("userID"))
	return c.JSON(input)
}

func listaractivos(c *fiber.Ctx) error {
	input := []gormdb.Suscripciones{}
	db.Find(&input, "Usuario_id = ? AND Activo = ?", c.Locals("userID"), true)
	return c.JSON(input)
}
func listarall(c *fiber.Ctx) error {
	input := []gormdb.Suscripciones{}
	db.Find(&input)
	return c.JSON(input)
}
