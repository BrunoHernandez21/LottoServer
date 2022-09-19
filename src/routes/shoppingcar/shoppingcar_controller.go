package shoppingcar

import (
	"lottomusic/src/models/compuestas"
	"lottomusic/src/models/gormdb"
	"time"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Carrito{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}

	if input.Plan_id == 0 {
		m["mensaje"] = "Id plna no puede ser null"
		return c.Status(500).JSON(m)
	}
	if input.Cantidad == 0 {
		m["mensaje"] = "Cantidad no puede ser null o 0"
		return c.Status(500).JSON(m)
	}
	id, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "error interno tokenError"
		return c.Status(500).JSON(m)
	}
	input.Usuario_id = id
	fecha := time.Now()
	input.Id = 0
	input.Fecha_carrito = &fecha
	activo := true
	input.Activo = &activo
	a := gormdb.Planes{}
	errdb := db.Table("planes").Find(&a, "id = ?", input.Plan_id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	amount := *a.Precio * float32(input.Cantidad)
	puntosLines := a.Puntos * input.Cantidad
	// Dinero
	input.Total_linea = &amount
	input.Precio_unitario = a.Precio
	// Puntos
	input.Puntos_linea = &puntosLines
	input.Puntos_unitario = &a.Puntos
	// moneda de intercambio
	input.Moneda = &a.Moneda
	if a.Suscribcion {
		db.
			Model(&gormdb.Carrito{}).
			Where("? = usuario_id AND plan_id IN (SELECT id from planes WHERE suscribcion = true)", id).
			Update("activo", false)
	}
	errdb = db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	errdb = db.Find(&input, "Usuario_id = ? AND fecha_carrito = ?", input.Usuario_id, input.Fecha_carrito)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}

func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Carrito{}
	errdb := db.Find(&a, "id = ? AND Usuario_id = ?", c.Params("id"), c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "El item no existe"
		return c.JSON(m)
	}
	errdb = db.Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}

func eliminarall(c *fiber.Ctx) error {
	m := make(map[string]string)
	errdb := db.Table("carrito").Where("Usuario_id = ?", c.Locals("userID")).Update("activo", false)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}

func listarWPlan(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	parse := []compuestas.CarritoPlan{}

	errdb := db.Table("carrito as c").
		Select(`c.id, c.cantidad, c.total_linea, c.puntos_linea, c.fecha_carrito, c.plan_id, p.titulo,p.descripcion, p.moneda, p.suscribcion`).
		Joins("INNER JOIN planes as p ON c.plan_id = p.id").
		Where("Usuario_id = ? AND c.Activo = ? ", c.Locals("userID"), true).
		Find(&parse)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["items_carrito"] = parse
	return c.JSON(m)
}

/*
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	id, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "error interno"
		return c.Status(500).JSON(m)
	}
	input := gormdb.Carrito{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	errdb := db.Find(&input, "id = ? AND Usuario_id = ?", input.Id, id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	if input.Id == 0 {
		m["mensaje"] = "El item no existe"
		return c.JSON(m)
	}

	errdb = db.Save(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	return c.JSON(input)
}
*/