package userevent

import (
	"lottomusic/src/models/gormdb"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func active_events(c *fiber.Ctx) error {

	resp := make(map[string]interface{})
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	input := []gormdb.Evento_usuario{}
	errdb := db.Find(&input, "Usuario_id = ? AND Activo = ?", c.Locals("userID"), true)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}
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

func winer_event(c *fiber.Ctx) error {
	resp := make(map[string]interface{})
	userID, err1 := c.Locals("userID").(uint32)
	if !err1 {
		resp["mensaje"] = "internal error"
		return c.Status(500).JSON(resp)
	}
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	a := int64(0)
	db.Table("ganador").Where("Usuario_id = ?", userID).Count(&a)
	pag, er4 := strconv.ParseUint(pagt, 10, 32)
	if er4 != nil {
		resp["mensaje"] = er4.Error()
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
	resp["pag"] = &pag
	resp["sizePage"] = &sizepage
	resp["totals"] = &a
	init := (pag - 1) * sizepage

	ganador := []gormdb.Ganador{}
	errdb := db.
		Table("ganador").
		Where("Usuario_id = ?", userID).
		Offset(int(init)).
		Limit(int(sizepage)).
		Find(&ganador)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["ganador"] = ganador

	return c.JSON(resp)
}

func history_event(c *fiber.Ctx) error {
	resp := make(map[string]interface{})
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	a := int64(0)
	userID := c.Locals("userID")
	db.Table("evento_usuario").Where("usuario_id = ? AND activo = ?", userID, false).Count(&a)
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
	apuestasUsuario := []gormdb.Evento_usuario{}
	errdb := db.
		Table("evento_usuario").
		Where("Usuario_id = ? AND activo = ?", userID, false).
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

func event_id(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Evento_usuario{}
	errdb := db.Find(&input, "Id = ?", c.Params("id"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}

func create_event(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Evento_usuario{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}

	if input.Evento_id == 0 {
		m["mensaje"] = "debe seleccionarse el evento"
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	// Validar el estado del ecento
	evento := gormdb.Eventos{}
	errdb := db.Find(&evento, "Id = ?", input.Evento_id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if !evento.Activo {
		m["mensaje"] = "Evento expirado"
		return c.Status(200).JSON(m)
	}

	// Cartera
	cartera := gormdb.Carteras{}
	errdb = db.Find(&cartera, "usuario_id = ?", userID)
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
	fecha := time.Now()
	input.Fecha = &fecha
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

// Root

// func userevent_list_all(c *fiber.Ctx) error {
// 	m := make(map[string]string)
// 	input := []gormdb.Evento_usuario{}
// 	errdb := db.Find(&input)
// 	if errdb.Error != nil {
// 		m["mensaje"] = errdb.Error.Error()
// 		return c.Status(500).JSON(m)
// 	}
// 	return c.JSON(input)
// }

func userevent_edit(c *fiber.Ctx) error {
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

func userevent_delete(c *fiber.Ctx) error {
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
