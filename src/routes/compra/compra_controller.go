package compra

import (
	"lottomusic/src/models/compuestas"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func verifica(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := inputs.Get_Stripe{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	var resp string
	db.Raw("CALL pagos_realizado(?,?)", input.Orden_id, input.StripeKey).Scan(&resp)
	m["resp"] = resp
	return c.Status(200).JSON(m)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Pagos{}
	err := db.Find(&a, "id = ?", param).Delete(&a)
	if err.Error != nil {
		m["mensaje"] = err.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {
	m := make(map[string]string)
	compra := []compuestas.Pagos_orden{}
	errdb := db.
		Table("pagos as p").
		Select(`p.id, p.fecha_pagado, p.orden_id, p.stripe_id, o.status, o.Fecha_emitido, o.Total, o.Iva, o.Descuento, o.Total_iva`).
		Joins("JOIN ordenes as o ON p.orden_id = o.id").
		Where("p.usuario_id = ?", c.Locals("userID")).
		Find(&compra)

	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(compra)
}

func listarpaginado(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	resp := make(map[string]interface{})

	a := int64(0)
	db.Table("pagos").Where("usuario_id = ?", userID).Count(&a)
	pag, err := strconv.ParseUint(c.Params("pag"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)
	if err != nil || err2 != nil {
		resp["mensaje"] = err.Error()
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

	compra := []compuestas.Pagos_orden{}
	errdb := db.
		Table("pagos as p").
		Select(`p.id, p.fecha_pagado, p.orden_id,	p.stripe_id, o.status, o.Fecha_emitido, o.Total, o.Iva, o.Descuento, o.Total_iva`).
		Joins("JOIN ordenes as o ON p.orden_id = o.id").
		Offset(int(init)).
		Limit(int(sizepage)).
		Where("p.usuario_id = ?", c.Locals("userID")).
		Find(&compra)

	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["compras"] = compra

	return c.JSON(resp)

}

func checkout(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	var resp string
	db.Raw("CALL genera_orden(?)", c.Locals("userID")).Scan(&resp)
	m["resp"] = resp
	return c.Status(200).JSON(m)
}

func createTarjeta(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Payment_method{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	input.Id = 0
	input.Usuario_id = userID
	input.Activo = true

	errdb := db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func editTarjeta(c *fiber.Ctx) error {

	m := make(map[string]string)
	input := gormdb.Payment_method{}
	compare := gormdb.Payment_method{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	input.Usuario_id = userID
	errdb := db.Find(&compare, "Id = ? ", input.Id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if compare.Id == 0 {
		m["mensaje"] = "No exite la tarjeta"
		return c.Status(500).JSON(m)
	}
	if compare.Id != input.Id {
		m["mensaje"] = "No te pertenece la tarjeta"
		return c.Status(500).JSON(m)
	}

	input.Usuario_id = userID
	input.Activo = true

	errdb = db.Save(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func deleteTarjeta(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Payment_method{}
	err := db.Find(&a, "id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Id == 0) {
		m["mensaje"] = "tarjeta no encontrado"
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	if a.Usuario_id != userID {
		m["mensaje"] = "Esta tarjeta no es tuya"
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
func listarTarjeta(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := []gormdb.Payment_method{}
	errdb := db.Find(&input, "usuario_id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["tarjetas"] = input
	return c.JSON(m)

}

func listarOrdenes(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Ordenes{}
	errdb := db.Find(&input, "usuario_id = ? AND status = ?", c.Locals("userID"), "proceso")
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)

}
