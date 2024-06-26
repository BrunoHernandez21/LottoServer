package apuesta

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Apuesta_usuario{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	cartera := gormdb.Carteras{}
	errdb := db.Find(&cartera, "Id_usuario = ?", userID)
	if errdb.Error != nil {
		return c.JSON(errdb)
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
		return c.JSON(errdb)
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
		return c.JSON(errdb)
	}
	errdb = db.Create(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}

	return c.JSON(input)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Apuesta_usuario{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if input.Id == 0 {
		m["mensaje"] = "Id no puede ser null"
		return c.Status(500).JSON(m)
	}
	db.Save(&input)
	return c.JSON(input)
}
func byid(c *fiber.Ctx) error {
	input := gormdb.Apuesta_usuario{}
	db.Find(&input, "Id = ?", c.Params("id"))
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)

	//db midelware
	a := gormdb.Apuesta_usuario{}
	err := db.Find(&a, "id = ?", c.Params("id")).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listarTodos(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Find(&input)
	return c.JSON(input)
}
func activosPage(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Find(&input, "Usuario_id = ? AND Activo = ?", c.Locals("userID"), true)
	return c.JSON(input)
}
func activo(c *fiber.Ctx) error {
	input := []gormdb.Apuesta_usuario{}
	db.Find(&input, "Usuario_id = ? AND Activo = ?", c.Locals("userID"), true)
	return c.JSON(input)
}
