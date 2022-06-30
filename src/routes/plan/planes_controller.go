package plan

import (
	"lottomusic/src/models/gormdb"

	"github.com/gofiber/fiber/v2"
)

func lista(c *fiber.Ctx) error {
	input := []gormdb.Plan{}
	db.Table("plan").Find(&input)
	return c.JSON(input)
}

func byname(c *fiber.Ctx) error {
	a := gormdb.Plan{}
	err2 := db.Table("plan").Find(&a, "nombre LIKE ?", "%"+c.Params("name")+"%")
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}
	return c.JSON(a)

}
func byid(c *fiber.Ctx) error {
	a := gormdb.Plan{}
	err2 := db.Table("plan").Find(&a, "id = ?", c.Params("id"))
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}
	return c.JSON(a)
}
func create(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Plan{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if (input.Precio == nil) || (input.Nombre == nil) || (input.Oportunidades == nil) {

		m["mensaje"] = "informacion insuficiente"
		return c.JSON(m)
	}

	input.Id = 0
	a := db.Table("plan").Create(&input)
	if a.Error != nil {

		m["mensaje"] = "No se pudo acceder a la base de datos"
		return c.JSON(m)
	}

	m["mensaje"] = "El plan ha sido creado"
	return c.JSON(m)
}
func delete(c *fiber.Ctx) error {

	m := make(map[string]string)
	a := gormdb.Plan{}
	err := db.Table("plan").Find(&a, "id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Id == 0) {
		m["mensaje"] = "Plan no encontrado"
		return c.JSON(m)
	}
	err2 := db.Table("plan").Delete(&a)
	if err2.Error != nil {
		m["mensaje"] = "No se pudo acceder a la base de datos"
		return c.JSON(m)
	}

	m["mensaje"] = "Eliminado con exito"
	return c.JSON(m)
}
func edit(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Plan{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if input.Id == 0 {
		m["mensaje"] = "informacion insuficiente"
		return c.JSON(m)
	}

	a := gormdb.Plan{}
	err2 := db.Table("plan").Find(&a, "id = ?", input.Id)
	if err2.Error != nil {
		m["mensaje"] = "No se pudo acceder a la base de datos"
		return c.JSON(m)
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

	err3 := db.Table("plan").Save(a)
	if err3.Error != nil {
		m["mensaje"] = "No se pudo acceder a la base de datos"
		return c.JSON(m)
	}

	m["mensaje"] = "Editado con exito"
	return c.JSON(m)
}
