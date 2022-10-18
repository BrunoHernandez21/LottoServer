package plan

import (
	"lottomusic/src/models/gormdb"
	"lottomusic/src/modules/stripe/impstripe"

	"github.com/gofiber/fiber/v2"
)

func single_payment(c *fiber.Ctx) error {
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

func list_subscriptions(c *fiber.Ctx) error {
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

	errdb := db.Table("planes").Find(&a, "titulo LIKE ?", "%"+c.Params("name")+"%")
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

///// root

func create(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Planes{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if (input.Precio == nil) || (input.Titulo == nil) || (input.Puntos == 0) {
		m["mensaje"] = "informacion insuficiente"
		return c.Status(500).JSON(m)
	}

	input.Id = 0
	errdb := db.Table("planes").Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	product, err := impstripe.Create_product(&input)
	if err != nil {
		db.Delete(&input)
		m["mensaje"] = "Stripe no esta disponible"
		return c.JSON(m)
	}
	price, err2 := impstripe.Create_price(&input, product.ID)
	if err2 != nil {
		db.Delete(&input)
		m["mensaje"] = "Stripe no esta disponible"
		return c.JSON(m)
	}
	input.Stripe_price = &price.ID
	input.Stripe_product = &product.ID
	db.Save(input)

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
	errdb := db.Find(&a, "id = ?", input.Id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if input.Puntos != 0 {
		a.Puntos = input.Puntos
	}
	if input.Precio != nil {
		a.Precio = input.Precio
	}
	if input.Titulo != nil {
		a.Titulo = input.Titulo
	}
	product, err := impstripe.Create_product(&a)
	if err != nil {
		m["mensaje"] = "Stripe no esta disponible"
		return c.JSON(m)
	}
	price, err2 := impstripe.Create_price(&a, product.ID)
	if err2 != nil {
		m["mensaje"] = "Stripe no esta disponible"
		return c.JSON(m)
	}
	a.Stripe_price = &price.ID
	a.Stripe_product = &product.ID
	errdb = db.Save(a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "Editado con exito"
	return c.JSON(m)
}
