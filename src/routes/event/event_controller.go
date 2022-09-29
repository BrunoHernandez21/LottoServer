package evento

import (
	"lottomusic/src/models/gormdb"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Eventos{}
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
	a := gormdb.Eventos{}
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

	input := gormdb.Eventos{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "Id es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Fechahora_evento == nil {
		m["mensaje"] = "Fechahora_evento es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Premio_cash == nil && input.Premio_otros == nil {
		m["mensaje"] = "Premio_cash o Premio_otros es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Premio_cash == nil && input.Moneda == nil {
		m["mensaje"] = "moneda no puede ser nul si el premio es cash"
		return c.Status(500).JSON(m)
	}
	if input.Video_id == 0 {
		m["mensaje"] = "video_id es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Costo == 0 {
		m["mensaje"] = "costo no puede ser nil"
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
	a := gormdb.Eventos{}
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

func listarActivos(c *fiber.Ctx) error {
	resp := make(map[string]interface{})
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	input := []gormdb.Eventos{}
	errdb := db.Find(&input, "Activo = ?", true)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}
	page, err4 := strconv.ParseUint(pagt, 0, 32)
	if err4 != nil {
		resp["mensaje"] = err4.Error()
		return c.Status(500).JSON(resp)
	}
	sizepage, err5 := strconv.ParseUint(sizepaget, 10, 32)
	if err5 != nil {
		resp["mensaje"] = err5.Error()
		return c.Status(500).JSON(resp)
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
