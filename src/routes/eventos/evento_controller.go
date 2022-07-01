package evento

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	input := gormdb.Apuestas{}
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
func byid(c *fiber.Ctx) error {
	param := c.Params("id")
	a := gormdb.Apuestas{}
	err2 := db.Find(&a, "id = ?", param)
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}
	if a.Id == 0 {
		m := make(map[string]string)
		m["mensaje"] = "El evento no existe"
		return c.JSON(m)
	}

	return c.JSON(a)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	m["mensaje"] = "Todas las entradas son nesesarias"
	input := gormdb.Apuestas{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if input.Id == 0 {

		return c.JSON(m)
	}

	if input.Fechahoraapuesta == nil {
		return c.JSON(m)
	}
	if input.Premio == nil {
		return c.JSON(m)
	}
	if input.Precio == nil {
		return c.JSON(m)
	}
	if input.Video_id == 0 {
		return c.JSON(m)
	}
	if input.Categoria_apuesta_id == 0 {
		return c.JSON(m)
	}
	if input.Tipo_apuesta_id == 0 {
		return c.JSON(m)
	}

	errdb := db.Save(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Apuestas{}
	err := db.Find(&a, "id = ?", param).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	if a.Id == 0 {
		m["mensaje"] = "El evento no existe"
		return c.JSON(m)
	}

	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listarTodos(c *fiber.Ctx) error {
	input := []gormdb.Apuestas{}
	db.Find(&input)
	return c.JSON(input)
}
func listarActivos(c *fiber.Ctx) error {
	input := []gormdb.Apuestas{}
	db.Find(&input, "activo = ?", true)
	return c.JSON(input)
}
func activo(c *fiber.Ctx) error {
	input := []gormdb.Apuestas{}
	db.Find(&input, "activo = ?", true)
	return c.JSON(input)
}
