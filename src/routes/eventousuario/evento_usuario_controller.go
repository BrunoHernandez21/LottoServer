package eventousuario

import (
	"lottomusic/src/models/gormdb"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Evento_usuario{}
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
	errdb := db.Find(&cartera, "usuario_id = ?", userID)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if cartera.Id == 0 {
		errdb = db.Create(&cartera)
		if errdb.Error != nil {
			m["mensaje"] = errdb.Error.Error()
			return c.Status(500).JSON(m)
		}
	}
	input.Id = 0
	activo := true
	input.Activo = &activo
	input.Usuario_id = userID
	if input.Evento_id == 0 {
		m["mensaje"] = "Apuesta id no puede ser nulo"
		return c.Status(400).JSON(m)
	}
	/// verificacion de apuesta
	var costo_unitario uint32
	db.
		Raw("SELECT costo from categoria_evento WHERE id = (SELECT categoria_evento_id from eventos WHERE id = ?);", input.Evento_id).
		Scan(&costo_unitario)
	var count uint32
	if input.Saved_count != nil {
		count += costo_unitario
	}
	if input.Views_count != nil {
		count += costo_unitario
	}
	if input.Like_count != nil {
		count += costo_unitario
	}
	if input.Dislikes_count != nil {
		count += costo_unitario
	}
	if input.Saved_count != nil {
		count += costo_unitario
	}

	if input.Comments_count != nil {
		count += costo_unitario
	}

	if count == 0 {
		m["mensaje"] = "la apuesta no puede estar vacia"
		return c.Status(400).JSON(m)
	}
	//// verificacion de saldo
	if count > cartera.Puntos {
		m["mensaje"] = "No tienes suficientes puntos"
		return c.Status(400).JSON(m)
	}
	cartera.Puntos -= count
	/// reducir en cartera

	errdb = db.Save(&cartera)
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
	input := gormdb.Evento_usuario{}
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
	input := gormdb.Evento_usuario{}
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
	a := gormdb.Evento_usuario{}
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
	input := []gormdb.Evento_usuario{}
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
	db.Table("evento_usuario").Where("Usuario_id = ?", userID).Count(&a)

	page, err := strconv.ParseUint(c.Params("page"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)

	if err != nil || err2 != nil {
		resp["mensaje"] = err.Error()
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
	apuestasUsuario := []gormdb.Evento_usuario{}
	errdb := db.
		Table("evento_usuario").
		Where("Usuario_id = ?", userID).
		Offset(int(init)).
		Limit(int(sizepage)).
		Find(&apuestasUsuario)
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
	input := []gormdb.Evento_usuario{}
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
