package apuesta

import (
	"lottomusic/src/models/gormdb"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Apuesta_usuario{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	cartera := gormdb.Carteras{}
	errdb := db.Find(&cartera, "Id_usuario = ?", userID)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	input.Id = 0
	activo := true
	input.Activo = &activo
	input.Cantidad = 0
	input.Usuario_id = userID
	if input.Apuesta_id == 0 {
		m["mensaje"] = "Apuesta id no puede ser nulo"
		return c.Status(400).JSON(m)
	}
	/// verificacion de apuesta
	if (input.Apuesta_id == 1) && (input.Vistas == 0) {
		m["mensaje"] = "Vistas no puede ser null"
		return c.Status(400).JSON(m)
	}
	if (input.Apuesta_id == 2) && (input.Likes == 0) {
		m["mensaje"] = "Vistas no puede ser null"
		return c.Status(400).JSON(m)
	}
	if (input.Apuesta_id == 3) && (input.Comentarios == 0) {
		m["mensaje"] = "Vistas no puede ser null"
		return c.Status(400).JSON(m)
	}
	//// verificacion de saldo
	evento := gormdb.Apuestas{
		Id: input.Apuesta_id,
	}
	errdb = db.Find(&evento)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	if (evento.Categoria_apuesta_id == 1) && (cartera.Oportunidades == 0) {
		m["mensaje"] = "No tienes de esta moneda"
		return c.Status(400).JSON(m)
	}
	if (evento.Categoria_apuesta_id == 2) && (cartera.Acumulado_alto8am == 0) {
		m["mensaje"] = "No tienes de esta moneda"
		return c.Status(400).JSON(m)
	}
	if (evento.Categoria_apuesta_id == 3) && (cartera.Acumulado_bajo8pm == 0) {
		m["mensaje"] = "No tienes de esta moneda"
		return c.Status(400).JSON(m)
	}
	if (evento.Categoria_apuesta_id == 4) && (cartera.Aproximacion_alta00am == 0) {
		m["mensaje"] = "No tienes de esta moneda"
		return c.Status(400).JSON(m)
	}
	if (evento.Categoria_apuesta_id == 5) && (cartera.Aproximacion_baja == 0) {
		m["mensaje"] = "No tienes de esta moneda"
		return c.Status(400).JSON(m)
	}
	/// reducir en cartera
	if evento.Categoria_apuesta_id == 1 {
		cartera.Oportunidades -= 1
	}
	if evento.Categoria_apuesta_id == 2 {
		cartera.Acumulado_alto8am -= 1
	}
	if evento.Categoria_apuesta_id == 3 {
		cartera.Acumulado_bajo8pm -= 1
	}
	if evento.Categoria_apuesta_id == 4 {
		cartera.Aproximacion_alta00am -= 1
	}
	if evento.Categoria_apuesta_id == 5 {
		cartera.Aproximacion_baja -= 1
	}

	errdb = db.Save(cartera)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	errdb = db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	return c.JSON(input)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Apuesta_usuario{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "Id no puede ser null"
		return c.Status(500).JSON(m)
	}
	db.Save(&input)
	return c.JSON(input)
}
func byid(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Apuesta_usuario{}
	errdb := db.Find(&input, "Id = ?", c.Params("id"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)

	//db midelware
	a := gormdb.Apuesta_usuario{}
	err := db.Find(&a, "id = ?", c.Params("id")).Delete(&a)
	if err.Error != nil {
		m["mensaje"] = err.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listarTodos(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Apuesta_usuario{}
	errdb := db.Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func historialPage(c *fiber.Ctx) error {

	resp := make(map[string]interface{})

	a := int64(0)
	userID := c.Locals("userID")
	db.Table("apuesta_usuario").Where("Usuario_id = ?", userID).Count(&a)

	page, err := strconv.ParseUint(c.Params("page"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)

	if err != nil || err2 != nil {
		resp["mensaje"] = err.Error()
		return c.Status(500).JSON(resp)
	}

	resp["pags"] = math.Round(float64(a) / float64(sizepage))
	resp["pag"] = &page
	resp["sizePage"] = &sizepage
	resp["totals"] = &a
	init := (page - 1) * sizepage
	apuestasUsuario := []gormdb.Apuesta_usuario{}
	errdb := db.Table("apuesta_usuario").Offset(int(init)).Limit(int(sizepage)).Find(&apuestasUsuario, "Usuario_id = ?", userID)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}
	resp["userEvent"] = apuestasUsuario
	return c.JSON(resp)
}

func activosPage(c *fiber.Ctx) error {
	m := make(map[string]string)
	resp := make(map[string]interface{})
	input := []gormdb.Apuesta_usuario{}
	errdb := db.Find(&input, "Usuario_id = ? AND Activo = ?", c.Locals("userID"), true)
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
func activo(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Apuesta_usuario{}
	errdb := db.Find(&input, "Usuario_id = ? AND Activo = ?", c.Locals("userID"), true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
