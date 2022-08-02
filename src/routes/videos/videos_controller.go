package videos

import (
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/views"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func videos_pag(c *fiber.Ctx) error {
	m := make(map[string]string)
	resp := make(map[string]interface{})
	input := []gormdb.Videos{}

	a := int64(0)
	db.Table("videos").Where("Activo = ?", true).Count(&a)
	page, err := strconv.ParseUint(c.Params("page"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)
	if err != nil || err2 != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	pags := uint64(a) / sizepage
	residuo := uint64(a) % sizepage
	if residuo != 0 {
		pags += 1
	}
	resp["pags"] = pags
	resp["pag"] = page
	resp["sizepage"] = sizepage
	resp["totals"] = a
	init := (page - 1) * sizepage
	errdb := db.
		Table("videos").
		Where("Activo = ?", true).
		Offset(int(init)).
		Limit(int(sizepage)).
		Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	resp["items"] = input
	return c.JSON(resp)
}

func videos_evento_pag(c *fiber.Ctx) error {
	m := make(map[string]string)
	resp := make(map[string]interface{})
	items := []views.EventoVideo{}

	a := int64(0)
	db.Find(&items).Count(&a)
	page, err := strconv.ParseUint(c.Params("page"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)
	if err != nil || err2 != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	pags := uint64(a) / sizepage
	residuo := uint64(a) % sizepage
	if residuo != 0 {
		pags += 1
	}
	resp["pags"] = pags
	resp["pag"] = page
	resp["sizePage"] = sizepage
	resp["totals"] = a
	init := (page - 1) * sizepage

	errdb := db.
		Table("eventos_videos").
		Offset(int(init)).
		Limit(int(sizepage)).
		Find(&items)

	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	resp["items"] = &items

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
	m := make(map[string]interface{})
	input := []string{}
	errdb := db.Table("videos").Select("genero").Where("Activo = ?", true).Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	input = uniqueString(input)
	m["grupos"] = input
	return c.JSON(m)
}

func listarGruposName(c *fiber.Ctx) error {
	m := make(map[string]string)
	resp := make(map[string]interface{})
	items := []views.EventoVideo{}
	genero := c.Params("name")
	a := int64(0)
	db.Find(&items).Where("genero = ?", genero).Count(&a)
	page, err := strconv.ParseUint(c.Params("page"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)
	if err != nil || err2 != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	pags := uint64(a) / sizepage
	residuo := uint64(a) % sizepage
	if residuo != 0 {
		pags += 1
	}
	resp["pags"] = pags
	resp["pag"] = page
	resp["sizePage"] = sizepage
	resp["totals"] = a
	init := (page - 1) * sizepage

	errdb := db.
		Table("eventos_videos").
		Where("genero = ?", genero).
		Offset(int(init)).
		Limit(int(sizepage)).
		Find(&items)

	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	resp["items"] = &items
	return c.JSON(resp)
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
	if input.Video_id != nil {
		ins.Video_id = input.Video_id
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

/*
func uniqueInt(arr []int) []int {
	result := make([]int, 0, len(arr))
	encountered := map[int]bool{}
	for v := range arr {
		encountered[arr[v]] = true
	}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}*/
