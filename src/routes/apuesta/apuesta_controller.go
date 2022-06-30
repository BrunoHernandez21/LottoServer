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
	input.Id = 0
	errdb := db.Create(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}

	return c.JSON(input)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Apuesta_usuario{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if input.Id == 0 {
		m["mensaje"] = "Id no puede ser null"
		return c.Status(500).JSON(m)
	}
	ins := gormdb.Apuesta_usuario{
		Id: input.Id,
	}
	err2 := db.Find(&ins)
	if err2.Error != nil {
		m["mensaje"] = "no existe"
		return c.JSON(m)
	}

	if input.Activo != nil {
		ins.Activo = input.Activo
	}
	if input.Apuesta != nil {
		ins.Apuesta = input.Apuesta
	}
	if input.Cantidad != 0 {
		ins.Cantidad = input.Cantidad
	}
	if input.Fecha != nil {
		ins.Fecha = input.Fecha
	}
	if input.Likes != nil {
		ins.Likes = input.Likes
	}
	if input.Monto != 0 {
		ins.Monto = input.Monto
	}
	if input.Suscribcion_id != nil {
		ins.Suscribcion_id = input.Suscribcion_id
	}
	if input.Usuario != nil {
		ins.Usuario = input.Usuario
	}
	if input.Vistas != nil {
		ins.Vistas = input.Vistas
	}

	db.Save(&ins)
	return c.JSON(ins)
}
func byid(c *fiber.Ctx) error {
	input := gormdb.Apuesta_usuario{}
	db.Find(&input, "Id = ?", c.Params("id"))
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)

	//db midelware
	a := gormdb.Apuesta_usuario{}
	err := db.Find(&a, "id = ?", c.Params("id")).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listarTodos(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Find(&input)
	return c.JSON(input)
}
func listarActivos(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Find(&input, "Usuario = ? AND Activo = ?", c.Locals("userID"), true)
	return c.JSON(input)
}
func activo(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Find(&input, "Usuario = ? AND Activo = ?", c.Locals("userID"), true)
	return c.JSON(input)
}
