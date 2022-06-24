package compra

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	input := gormdb.Compra{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	errdb := db.Model("Compra").Create(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}
	input.Id = 0
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Compra{}
	err := db.Model("Compra").Find(&a, "id = ?", param).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	m["mensjae"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {
	input := []gormdb.Compra{}
	db.Model("Compra").Find(&input)
	return c.JSON(input)
}
func checkout(c *fiber.Ctx) error {
	m := make(map[string]string)
	m["mensaje"] = "Todas las entradas son nesesarias"
	input := gormdb.Orden{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if input.Id == 0 {
		return c.JSON(m)
	}
	db.Model("Orden").Find(&input)
	order := "En comprovacion"
	input.Orden_status = &order
	errdb := db.Model("Orden").Save(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}
	return c.JSON(input)
}
