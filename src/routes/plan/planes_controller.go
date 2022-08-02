package plan

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func listar_one(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := []gormdb.Planes{}
	errdb := db.Table("plan_one").Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["planes"] = input
	return c.JSON(m)
}

func lista_suscripcion(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := []gormdb.Planes{}
	errdb := db.Table("plan_suscripcion").Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["planes"] = input
	return c.JSON(m)
}

func byname(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Planes{}
	errdb := db.Table("planes").Find(&a, "nombre LIKE ?", "%"+c.Params("name")+"%")
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(a)

}
func byid(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Planes{}
	errdb := db.Table("planes").Find(&a, "id = ?", c.Params("id"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(a)
}
func create(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Planes{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if (input.Precio == nil) || (input.Nombre == nil) || (input.Cash == nil) {
		m["mensaje"] = "informacion insuficiente"
		return c.Status(500).JSON(m)
	}

	input.Id = 0
	errdb := db.Table("planes").Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "El plan ha sido creado"
	return c.JSON(m)
}
func delete(c *fiber.Ctx) error {

	m := make(map[string]string)
	a := gormdb.Planes{}
	err := db.Find(&a, "id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Id == 0) {
		m["mensaje"] = "Plan no encontrado"
		return c.Status(500).JSON(m)
	}
	errdb := db.Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "Eliminado con exito"
	return c.JSON(m)
}
func edit(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Planes{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "informacion insuficiente"
		return c.JSON(m)
	}

	a := gormdb.Planes{}
	errdb := db.Table("planes").Find(&a, "id = ?", input.Id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if input.Cash != nil {
		a.Cash = input.Cash
	}
	if input.Precio != nil {
		a.Precio = input.Precio
	}
	if input.Nombre != nil {
		a.Nombre = input.Nombre
	}

	errdb = db.Table("plplanesan").Save(a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "Editado con exito"
	return c.JSON(m)
}
