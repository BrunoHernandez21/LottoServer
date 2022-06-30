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
		return err
	}

	if input.Id_plan == nil {
		m["mensaje"] = "Id plna no puede ser null"
		return c.JSON(m)
	}
	if input.Cantidad == 0 {
		m["mensaje"] = "Cantidad no puede ser null o 0"
		return c.JSON(m)
	}

	if input.Cantidad == 0 {
		m["mensaje"] = "Cantidad no puede ser null o 0"
		return c.JSON(m)
	}

	id, ok := c.Locals("userID").(uint32)
	if ok {
		input.Usuario_id = &id
	} else {
		m["mensaje"] = "error interno"
		return c.JSON(m)
	}
	status := "Carrito"
	fecha := time.Now()
	input.Id = 0
	input.Orden_status = &status
	input.Fecha_orden = &fecha

	errdb := db.Create(&input)
	if errdb.Error != nil {
		return c.JSON(errdb.Error)
	}
	errdb = db.Find(&input, "Usuario_id = ? AND Fecha_orden = ?", input.Usuario_id, input.Fecha_orden)
	if errdb.Error != nil {
		return c.JSON(errdb.Error)
	}
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Orden{}
	err := db.Find(&a, "id = ?", param).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}

	if a.Id == 0 {
		m["mensaje"] = "El item no existe"
	} else {
		m["mensaje"] = "Eliminado Satisfactoriamente"
	}

	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {
	input := []gormdb.Orden{}
	db.Find(&input, "Usuario_id = ?", c.Locals("userID"))
	return c.JSON(input)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Orden{}
	if err := c.BodyParser(&input); err != nil {
		return err
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

	er := db.Save(out)
	if er.Error != nil {
		return c.JSON(er.Error)
	}
	return c.JSON(out)
}
