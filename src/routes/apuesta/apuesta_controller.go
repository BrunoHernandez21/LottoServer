package apuesta

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	input := gormdb.Apuesta_usuario{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	errdb := db.Model("Apuesta_usuario").Create(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}
	input.Id = 0
	return c.JSON(input)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", c.Params("id"))
	if err2.Error != nil {
		m["mensjae"] = "Usuario no registrado"
		return c.JSON(m)
	}
	input := gormdb.Usuarios{}
	db.Model("Apuesta_usuario").Save(&input)
	return c.JSON(input)
}
func byid(c *fiber.Ctx) error {
	input := gormdb.Apuesta_usuario{}
	db.Model("Apuesta_usuario").Find(&input, "Id = ?", c.Params("id"))
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)

	//db midelware
	a := gormdb.Apuesta_usuario{}
	err := db.Model("Apuesta_usuario").Find(&a, "id = ?", c.Params("id")).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	m["mensjae"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listarTodos(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Model("Apuesta_usuario").Find(&input)
	return c.JSON(input)
}
func listarActivos(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Model("Apuesta_usuario").Find(&input, "Usuario = ? AND Activo = ?", c.Locals("userID"), true)
	return c.JSON(input)
}
func activo(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Model("Apuesta_usuario").Find(&input, "Usuario = ? AND Activo = ?", c.Locals("userID"), true)
	return c.JSON(input)
}
