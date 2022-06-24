package carrito

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	input := gormdb.Orden{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	errdb := db.Model("Orden").Create(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}
	input.Id = 0
	db.Model("Orden").Save(&input)
	db.Model("Orden").Find(&input, "Usuario_id = ?", c.Locals("userID")).Last(&input)
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Orden{}
	err := db.Model("Orden").Find(&a, "id = ?", param).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	m["mensjae"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {
	input := []gormdb.Orden{}
	db.Model("Orden").Find(&input, "Usuario_id = ?", c.Locals("userID"))
	return c.JSON(input)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Orden{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if input.Id == 0 {
		m["mensaje"] = "El id es necesario"
		return c.JSON(m)
	}
	db.Model("Orden").Find(&input)
	return c.JSON(input)
}
