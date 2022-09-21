package videos

import (
	"lottomusic/src/helpers"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/views"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//////// Eventos video paginado
func videos_evento_pag(c *fiber.Ctx) error {
	m := make(map[string]string)
	resp := make(map[string]interface{})
	items := []views.EventoVideo{}
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	a := int64(0)
	db.Find(&items).Count(&a)
	page, err := strconv.ParseUint(pagt, 10, 32)
	if err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	sizepage, err2 := strconv.ParseUint(sizepaget, 10, 32)
	if err2 != nil {
		m["mensaje"] = err2.Error()
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

///////// Videos
func videos_pag(c *fiber.Ctx) error {
	m := make(map[string]string)
	resp := make(map[string]interface{})
	input := []gormdb.Videos{}
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	a := int64(0)
	db.Table("videos").Where("Activo = ?", true).Count(&a)
	page, err := strconv.ParseUint(pagt, 10, 32)
	if err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	sizepage, err2 := strconv.ParseUint(sizepaget, 10, 32)
	if err2 != nil {
		m["mensaje"] = err2.Error()
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

func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Videos{}
	err := db.Find(&a, "id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Id == 0) {
		m["mensaje"] = "Video no encontrado"
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

////////////// Grupos
func listargrupos(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := []string{}
	errdb := db.Table("videos").Select("genero").Where("Activo = ?", true).Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	input = helpers.UniqueString(input)
	m["grupos"] = input
	return c.JSON(m)
}

func listarGruposName(c *fiber.Ctx) error {

	resp := make(map[string]interface{})

	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	genero := c.Params("name")
	items := []views.EventoVideo{}
	a := int64(0)
	db.Find(&items).Where("genero = ?", genero).Count(&a)
	page, err4 := strconv.ParseUint(pagt, 10, 32)
	if err4 != nil {
		resp["mensaje"] = err4.Error()
		return c.Status(500).JSON(resp)
	}
	sizepage, err5 := strconv.ParseUint(sizepaget, 10, 32)
	if err5 != nil {
		resp["mensaje"] = err5.Error()
		return c.Status(500).JSON(resp)
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
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["items"] = &items
	return c.JSON(resp)
}

/////////////// Estadisticas

func get_statistics(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := []gormdb.VideosEstadisticas{}

	err := db.Order("Video_id DESC").Find(&input)
	if err.Error != nil {
		m["mensaje"] = err.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["items"] = input
	return c.JSON(m)
}
func get_st_byID(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.VideosEstadisticas{}

	err := db.Where("Video_id = ?", c.Params("id")).Order("Video_id ASC").Find(&input)
	if err.Error != nil {
		m["mensaje"] = err.Error.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "No hay estadisticas del video"
		return c.JSON(m)
	}
	return c.JSON(input)
}

//TODO incompleto
func delete_statistics(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.VideosEstadisticas{}
	err := db.Find(&a, "Video_id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Video_id == 0) {
		m["mensaje"] = "Video no encontrado"
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

func edit_statistics(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.VideosEstadisticas{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Video_id == 0 {
		m["mensaje"] = "Id no puede ser null"
		return c.Status(500).JSON(m)
	}

	errdb := db.Table("videos_estadisticas").Where("Video_id = ?", input.Video_id).Save(&input)

	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}

// func create_statistics(c *fiber.Ctx) error {
// 	m := make(map[string]string)
// 	input := gormdb.VideosEstadisticas{}
// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(500).JSON(err)
// 	}
// 	if input.Video_id == 0 {
// 		m["mensaje"] = "requieres incertar el video ID"
// 		return c.Status(400).JSON(m)
// 	}
// 	fecha := time.Now()
// 	input.Fecha = fecha
// 	errdb := db.Create(&input)
// 	if errdb.Error != nil {
// 		m["mensaje"] = errdb.Error.Error()
// 		return c.Status(500).JSON(m)
// 	}
// 	m["mensaje"] = "Creado con exito"
// 	m["fecha"] = fecha.String()
// 	return c.JSON(time.Now())
// }
