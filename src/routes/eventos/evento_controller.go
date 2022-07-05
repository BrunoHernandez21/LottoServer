package evento

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Apuestas{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	input.Id = 0
	errdb := db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func byid(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	a := gormdb.Apuestas{}
	errdb := db.Find(&a, "id = ?", param)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "El evento no existe"
		return c.Status(500).JSON(m)
	}

	return c.JSON(a)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	m["mensaje"] = "Todas las entradas son nesesarias"
	input := gormdb.Apuestas{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		return c.Status(500).JSON(m)
	}
	if input.Fechahoraapuesta == nil {
		return c.Status(500).JSON(m)
	}
	if input.Premio == nil {
		return c.Status(500).JSON(m)
	}
	if input.Precio == nil {
		return c.Status(500).JSON(m)
	}
	if input.Video_id == 0 {
		return c.Status(500).JSON(m)
	}
	if input.Categoria_apuesta_id == 0 {
		return c.Status(500).JSON(m)
	}
	if input.Tipo_apuesta_id == 0 {
		return c.Status(500).JSON(m)
	}

	errdb := db.Save(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Apuestas{}
	errdb := db.Find(&a, "id = ?", param).Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "El evento no existe"
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listarTodos(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Apuestas{}
	errdb := db.Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func listarActivos(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Apuestas{}
	errdb := db.Find(&input, "activo = ?", true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func activo(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Apuestas{}
	errdb := db.Find(&input, "activo = ?", true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
