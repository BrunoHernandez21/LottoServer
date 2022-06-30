package compra

import (
	"lottomusic/src/models/compra"
	"lottomusic/src/models/gormdb"
	"time"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := compra.Get_Checkout{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	orders := []gormdb.Orden{}
	errdb := db.Where("id IN ?", input.IDs).Find(&orders)
	if errdb.Error != nil {
		m["mensjae"] = "Internal Error"
		return c.Status(500).JSON(m)
	}
	for _, order := range orders {
		//retirar del carrtio
		activo := false
		order.Activa = &activo
		finalizado := "finalizado"
		order.Orden_status = &finalizado
		errdb := db.Save(&order)
		if errdb.Error != nil {
			m["mensaje"] = "Error interno"
			return c.Status(500).JSON(m)
		}
		//update cartera
		cartera := gormdb.Carteras{
			Id: order.Usuario_id,
		}
		errdb = db.Find(&cartera)
		if errdb.Error != nil {
			m["mensaje"] = "Error interno"
			return c.Status(500).JSON(m)
		}
		plan := gormdb.Plan{
			Id: order.Id_plan,
		}
		errdb = db.Find(&plan)
		if errdb.Error != nil {
			m["mensaje"] = "Error interno"
			return c.Status(500).JSON(m)
		}
		cartera.Acumulado_alto8am += *plan.Acumulado_alto8am
		cartera.Acumulado_bajo8pm += *plan.Acumulado_bajo8pm
		cartera.Aproximacion_alta00am += *plan.Aproximacion_alta00am
		cartera.Aproximacion_baja += *plan.Aproximacion_baja
		cartera.Oportunidades += *plan.Oportunidades

		db.Save(&cartera)
		//ADD al registro de compras
		tiempo := time.Now()
		compra := gormdb.Compra{
			Id:           0,
			Cantidad:     order.Cantidad,
			Amount:       *order.Amount,
			Fecha_compra: &tiempo,
			Usuario_id:   order.Usuario_id,
			Plan_id:      order.Id_plan,
		}
		errdb = db.Create(&compra)
		if errdb.Error != nil {
			m["mensaje"] = "Error interno"
			return c.Status(500).JSON(m)
		}
	}
	m["mensaje"] = "ok"
	return c.JSON(m)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Compra{}
	err := db.Find(&a, "id = ?", param).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {
	input := []gormdb.Compra{}
	db.Find(&input, "Usuario_id = ?", c.Locals("userID"))
	return c.JSON(input)
}
func checkout(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := compra.Get_Checkout{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if len(input.IDs) == 0 {
		m["mensaje"] = "it's empty"
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}

	orders := []gormdb.Orden{}
	errdb := db.Where("id IN ?", input.IDs).Find(&orders)
	if errdb.Error != nil {
		m["mensjae"] = "Internal Error"
		return c.Status(500).JSON(m)
	}
	for _, order := range orders {
		if !*order.Activa {
			m["mensaje"] = "Esta orden ha expirado"
			m["OrderID"] = order.Id
			return c.Status(500).JSON(m)
		}
		if order.Usuario_id != userID {
			m["mensaje"] = "Esta orden no te pertenece"
			m["OrderID"] = order.Id
			return c.Status(500).JSON(m)
		}
		if *order.Orden_status == "comprobación" {
			m["mensaje"] = "Esta compra esta en proceso"
			return c.Status(500).JSON(m)
		}
		comprobacion := "comprobación"
		order.Orden_status = &comprobacion
		errdb := db.Save(&order)
		if errdb.Error != nil {
			m["mensaje"] = "Error interno"
			return c.Status(500).JSON(m)
		}
	}
	c.App().Post("", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.GetRespHeader("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNjU2Njk5MzU3fQ.6kpTcYFL22Q-UdhsHLvLkSY7qQVp-vWaZcTq3HouU2QJja-0qs5xexf182NdrsGTUjXO4rghPfkx2YjNqmbH6g")
		if err := c.BodyParser(input); err != nil {
			return err
		}
		return nil
	})
	m["mensaje"] = "su compra esta en Comprobacion"
	return c.JSON(input)
}
