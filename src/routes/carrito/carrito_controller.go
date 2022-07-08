package carrito

import (
	"lottomusic/src/models/gormdb"
	"time"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Orden{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}

	if input.Id_plan == 0 {
		m["mensaje"] = "Id plna no puede ser null"
		return c.Status(500).JSON(m)
	}
	if input.Cantidad == 0 {
		m["mensaje"] = "Cantidad no puede ser null o 0"
		return c.Status(500).JSON(m)
	}

	id, ok := c.Locals("userID").(uint32)
	if ok {
		input.Usuario_id = id
	} else {
		m["mensaje"] = "error interno"
		return c.Status(500).JSON(m)
	}
	status := "Carrito"
	fecha := time.Now()
	input.Id = 0
	input.Orden_status = &status
	input.Fecha_orden = &fecha
	activo := true
	input.Activa = &activo
	a := gormdb.Plan{}
	errdb := db.Table("plan").Find(&a, "id = ?", input.Id_plan)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	amount := *a.Precio * float32(input.Cantidad)
	input.Amount = &amount

	errdb = db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	errdb = db.Find(&input, "Usuario_id = ? AND Fecha_orden = ?", input.Usuario_id, input.Fecha_orden)
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
	a := gormdb.Orden{}
	errdb := db.Find(&a, "id = ?", param).Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	if a.Id == 0 {
		m["mensaje"] = "El item no existe"
	} else {
		m["mensaje"] = "Eliminado Satisfactoriamente"
	}

	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Orden{}
	errdb := db.Find(&input, "Usuario_id = ? AND Activa = ? ", c.Locals("userID"), true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func listarWPlan(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	ordenes := []gormdb.Orden{}
	errdb := db.Find(&ordenes, "Usuario_id = ? AND Activa = ? ", c.Locals("userID"), true)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	result := make([]uint32, 0, len(ordenes))
	encountered := map[uint32]bool{}
	for v := range ordenes {
		encountered[ordenes[v].Id_plan] = true
	}
	for key := range encountered {
		result = append(result, key)
	}

	planes := []gormdb.Plan{}
	errdb = db.Find(&planes, "Id IN ? ", result)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	rp := make([]map[string]interface{}, 0, len(ordenes))
	for _, a := range ordenes {
		for _, v := range planes {
			if a.Id_plan == v.Id {
				mapa := make(map[string]interface{})
				mapa["plane"] = v
				mapa["orden"] = a
				rp = append(rp, mapa)
			}
		}
	}
	m["Ordenes"] = rp
	return c.JSON(m)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Orden{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "El id es necesario"
		return c.JSON(m)
	}
	out := gormdb.Orden{
		Id: input.Id,
	}
	db.Find(&out)
	out.Cantidad = input.Cantidad
	out.Id_plan = input.Id_plan
	out.Id_charges = input.Id_charges
	out.Amount = input.Amount

	errdb := db.Save(out)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(out)
}
