package plan

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func lista(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := []gormdb.Plan{}
	errdb := db.Table("plan").Find(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = nil
	m["planes"] = input
	return c.JSON(m)
}

func byname(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Plan{}
	errdb := db.Table("plan").Find(&a, "nombre LIKE ?", "%"+c.Params("name")+"%")
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(a)

}
func byid(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Plan{}
	errdb := db.Table("plan").Find(&a, "id = ?", c.Params("id"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(a)
}
func create(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Plan{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if (input.Precio == nil) || (input.Nombre == nil) || (input.Oportunidades == nil) {
		m["mensaje"] = "informacion insuficiente"
		return c.Status(500).JSON(m)
	}

	input.Id = 0
	errdb := db.Table("plan").Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "El plan ha sido creado"
	return c.JSON(m)
}
func delete(c *fiber.Ctx) error {

	m := make(map[string]string)
	a := gormdb.Plan{}
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
	input := gormdb.Plan{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "informacion insuficiente"
		return c.JSON(m)
	}

	a := gormdb.Plan{}
	errdb := db.Table("plan").Find(&a, "id = ?", input.Id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if input.Acumulado_alto8am != nil {
		a.Acumulado_alto8am = input.Acumulado_alto8am
	}
	if input.Acumulado_bajo8pm != nil {
		a.Acumulado_bajo8pm = input.Acumulado_bajo8pm
	}
	if input.Aproximacion_alta00am != nil {
		a.Aproximacion_alta00am = input.Aproximacion_alta00am
	}
	if input.Aproximacion_baja != nil {
		a.Aproximacion_baja = input.Aproximacion_baja
	}
	if input.Precio != nil {
		a.Precio = input.Precio
	}
	if input.Nombre != nil {
		a.Nombre = input.Nombre
	}
	if input.Oportunidades != nil {
		a.Oportunidades = input.Oportunidades
	}

	errdb = db.Table("plan").Save(a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "Editado con exito"
	return c.JSON(m)
}
