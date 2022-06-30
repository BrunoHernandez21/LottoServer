package videos

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func listar(c *fiber.Ctx) error {
	input := []gormdb.Videos{}
	db.Find(&input)
	return c.JSON(input)
}

func listaractivos(c *fiber.Ctx) error {
	input := []gormdb.Videos{}
	db.Find(&input, "Activo = ?", true)
	return c.JSON(input)
}

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Videos{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(err)
	}
	input.Id = 0
	errdb := db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = "Error en la base de datos"
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Creado con exito"
	return c.JSON(m)
}

func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Videos{}
	err := db.Find(&a, "id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Id == 0) {
		m["mensaje"] = "Plan no encontrado"
		return c.Status(404).JSON(m)
	}
	err2 := db.Delete(&a)
	if err2.Error != nil {
		m["mensaje"] = "No se pudo acceder a la base de datos"
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "eliminado correctamente"
	return c.JSON(m)
}

func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Videos{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if input.Id == 0 {
		m["mensaje"] = "Id no puede ser null"
		return c.Status(500).JSON(m)
	}
	ins := gormdb.Videos{
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
	if input.Artista != nil {
		ins.Artista = input.Artista
	}
	if input.Canal != nil {
		ins.Canal = input.Canal
	}
	if input.Fecha_video != nil {
		ins.Fecha_video = input.Fecha_video
	}
	if input.Id_video != nil {
		ins.Id_video = input.Id_video
	}
	if input.Titulo != nil {
		ins.Titulo = input.Titulo
	}
	if input.Url_video != nil {
		ins.Url_video = input.Url_video
	}

	db.Save(&ins)
	return c.JSON(ins)
}
