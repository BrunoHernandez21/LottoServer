package videos

import (
	"lottomusic/src/models/gormdb"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func listar(c *fiber.Ctx) error {
	input := []gormdb.Videos{}
	db.Find(&input)
	return c.JSON(input)
}

func listaractivos(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Videos{}
	errdb := db.Find(&input, "Activo = ?", true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func pagelistar(c *fiber.Ctx) error {
	m := make(map[string]string)
	resp := make(map[string]interface{})
	input := []gormdb.Videos{}
	errdb := db.Find(&input, "Activo = ?", true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	page, err := strconv.ParseUint(c.Params("page"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)
	if err != nil || err2 != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}

	resp["pags"] = math.Round(float64(len(input)) / float64(sizepage))
	resp["pag"] = page
	init := (page - 1) * sizepage
	end := (page * sizepage) - 1
	if int(end) > len(input) {
		end = uint64(len(input))
	}
	if init > end {
		resp["videos"] = nil
	} else {
		resp["videos"] = input[init:end]
	}

	return c.JSON(resp)
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
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Creado con exito"
	return c.JSON(m)
}
func listargrupos(c *fiber.Ctx) error {
	input := []string{}
	db.Table("Videos").Select("genero").Where("Activo = ?", true).Find(&input)
	input = uniqueString(input)
	return c.JSON(input)
}

func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Videos{}
	err := db.Find(&a, "id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Id == 0) {
		m["mensaje"] = "Plan no encontrado"
		return c.Status(404).JSON(m)
	}
	errdb := db.Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = "No se pudo acceder a la base de datos"
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "eliminado correctamente"
	return c.JSON(m)
}

func activoID(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Videos{}
	errdb := db.Find(&input, "id = ? AND Activo = ?", c.Params("id"), true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}

func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Videos{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "Id no puede ser null"
		return c.Status(500).JSON(m)
	}
	ins := gormdb.Videos{
		Id: input.Id,
	}
	errdb := db.Find(&ins)
	if errdb.Error != nil {
		m["mensaje"] = "no existe"
		return c.Status(500).JSON(m)
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

	errdb = db.Save(&ins)

	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(ins)
}

func uniqueString(arr []string) []string {
	result := make([]string, 0, len(arr))
	encountered := map[string]bool{}
	for v := range arr {
		encountered[arr[v]] = true
	}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
