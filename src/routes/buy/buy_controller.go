package buy

import (
	"encoding/json"
	"lottomusic/src/models/compuestas"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"lottomusic/src/modules/stripe/stripeme"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

////// Historial

func buy_history_paginated(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	resp := make(map[string]interface{})
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	a := int64(0)
	db.
		Table("pagos_orden").
		Where("usuario_id = ? AND status = ?", userID, "pagado").
		Count(&a)
	pag, err4 := strconv.ParseUint(pagt, 10, 32)
	if err4 != nil {
		resp["mensaje"] = err4.Error()
		return c.Status(500).JSON(resp)
	}
	sizepage, err5 := strconv.ParseUint(sizepaget, 10, 32)
	if err5 != nil {
		resp["mensaje"] = err5.Error()
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
		Table("pagos_orden").
		Offset(int(init)).
		Limit(int(sizepage)).
		Where("usuario_id = ? AND status = ?", c.Locals("userID"), "pagado").
		Find(&compra)

	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["compras"] = compra

	return c.JSON(resp)

}

func list_orders(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	resp := make(map[string]interface{})

	compra := []compuestas.Pagos_orden{}
	errdb := db.Table("pagos_orden").Where("usuario_id = ? AND status = ?", userID, "proceso").Find(&compra)

	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["compras"] = compra

	return c.JSON(resp)

}

func list_orders_errors(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	resp := make(map[string]interface{})

	compra := []compuestas.Pagos_orden{}
	errdb := db.Table("pagos_orden").Where("usuario_id = ? AND status = ?", userID, "rechazado").Find(&compra)

	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["compras"] = compra

	return c.JSON(resp)

}

func create_order(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	//generar y obtener la orden
	orden := gormdb.Ordenes{}
	db.Raw("CALL genera_orden_unico(?)", c.Locals("userID")).Scan(&orden)
	if orden.Id == 0 {
		m["mensaje"] = "Carrito vacio"
		return c.Status(500).JSON(m)
	}

	items_orden := []gormdb.ItemsOrden{}
	db.Find(&items_orden, "Orden_id = ?", orden.Id)
	m["orden"] = orden
	m["items_orden"] = items_orden
	return c.Status(500).JSON(m)
}

func create_payment_intent(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := inputs.GenerarPaymentItent{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "error al parcear datos de entrada"
		return c.Status(500).JSON(m)
	}
	if input.Orden_id == 0 {
		m["mensaje"] = "orden no puede estar vacio"
		return c.Status(500).JSON(m)
	}
	orden := gormdb.Ordenes{}
	db.Find(&orden, "id = ? AND status = ?", input.Orden_id, "proceso")
	if orden.Id == 0 {
		m["mensaje"] = "La orden expiro o no existe"
		return c.Status(500).JSON(m)
	}
	//generar y obtener la orden
	a, err := stripeme.Create_payment_intent(&orden)
	if err != nil {
		m["mensaje"] = "Stripe error"
		return c.Status(500).JSON(m)
	}
	m["id"] = a.ID
	m["status"] = a.Status
	m["amount"] = a.Amount
	m["client_secret"] = a.ClientSecret
	return c.JSON(m)
}

// checkout
func checkout(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})
	input := inputs.Checkout{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "error al parcear datos de entrada"
		return c.Status(500).JSON(m)
	}
	if input.Orden_id == 0 || input.Stripe_Payment == "" {
		m["mensaje"] = "Card y stripe_payment no pueden ser nulo"
		return c.Status(500).JSON(m)
	}

	//obtener la orden
	orden := gormdb.Ordenes{}
	db.Find(&orden, "id = ?", input.Orden_id)
	if orden.Id == 0 {
		m["mensaje"] = "Carrito vacio"
		return c.Status(500).JSON(m)
	}

	// mandamos a stripe a generar el intento de pago
	resp, errstr := stripeme.Pay_payment_intent(&orden, input.Stripe_Payment)
	var outReason string
	if errstr != nil {
		// Compra fallida
		db.Raw("CALL orden_rechazada(?,?)", orden.Id, errstr.Error()).Scan(&outReason)
		m["mensaje"] = errstr.Error()
		return c.Status(200).JSON(m)
	}
	// Compra realizada
	data, err := json.Marshal(&resp)
	if err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	db.Raw("CALL orden_pagada(?,?)", orden.Id, string(data)).Scan(&outReason)
	m["resp"] = "Compra realizada con éxito"
	return c.Status(200).JSON(m)
}

// checkout
func subscription_orden(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	input := inputs.SuscripcionOrden{}
	orden := gormdb.Ordenes{}
	m := make(map[string]interface{})

	db.Raw("CALL genera_orden_unico(?,?)", c.Locals("userID"), input.Plan_id).Scan(&orden)
	if orden.Id == 0 {
		m["mensaje"] = "Error interno"
		return c.Status(500).JSON(m)
	}

	return c.Status(200).JSON(orden)
}

func subscription_checkout(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})
	m["resp"] = "Compra realizada con éxito"
	return c.Status(200).JSON(m)
}

// func buy_retry(c *fiber.Ctx) error {
// 	/// Verificar la respuesta del usuario
// 	m := make(map[string]interface{})
// 	input := inputs.RetryCheckout{}
// 	if err := c.BodyParser(&input); err != nil {
// 		m["mensaje"] = "error al parcear datos de entrada"
// 		return c.Status(500).JSON(m)
// 	}
// 	if input.Card_id == 0 {
// 		m["mensaje"] = "Card no puede ser nulo o 0"
// 		return c.Status(500).JSON(m)
// 	}

// 	// obtenemos la tarjeta del usuario
// 	userID, isit := c.Locals("userID").(uint32)
// 	if !isit {
// 		m["mensaje"] = "internal error"
// 		return c.Status(500).JSON(m)
// 	}
// 	shaccv := mysha.SHA256(input.Cvc)
// 	tarjeta := gormdb.Payment_method{}
// 	errdb := db.Find(&tarjeta, "Id = ? ", input.Card_id)
// 	if errdb.Error != nil {
// 		m["mensaje"] = errdb.Error.Error()
// 		return c.Status(500).JSON(m)
// 	}
// 	if tarjeta.Usuario_id != userID {
// 		m["mensaje"] = "La tarjeta no te pertenece"
// 		return c.Status(500).JSON(m)
// 	}
// 	if tarjeta.Cvc != shaccv {
// 		m["mensaje"] = "Error de CCV"
// 		return c.Status(500).JSON(m)
// 	}

// 	// obtener la orden
// 	orden := gormdb.Ordenes{}
// 	errdb = db.Find(&orden, "Id = ? ", input.Orden_id)
// 	if errdb.Error != nil {
// 		m["mensaje"] = errdb.Error.Error()
// 		return c.Status(500).JSON(m)
// 	}
// 	if orden.Usuario_id != userID {
// 		m["mensaje"] = "La tarjeta no te pertenece"
// 		return c.Status(500).JSON(m)
// 	}

// 	// mandamos a stripe a verificar la compra
// 	resp, errstr := stripeme.Payment(tarjeta.ToStripeMethod(input.Cvc), &orden)
// 	var outReason string
// 	if errstr != nil {
// 		//Compra fallida
// 		db.Raw("CALL pagos_rechazado(?,?)", orden.Id, errstr.Error()).Scan(&outReason)
// 		m["mensaje"] = errstr.Error()
// 		return c.Status(200).JSON(m)
// 	}

// 	data, err := json.Marshal(&resp)
// 	if err != nil {
// 		m["mensaje"] = err.Error()
// 		return c.Status(500).JSON(m)
// 	}

// 	db.Raw("CALL pago_unico(?,?)", orden.Id, string(data)).Scan(&outReason)
// 	return c.Status(200).JSON(outReason)
// }

// func buy_cancel(c *fiber.Ctx) error {
// 	m := make(map[string]interface{})
// 	input := inputs.RetryCheckout{}
// 	if err := c.BodyParser(&input); err != nil {
// 		m["mensaje"] = "error al parcear datos de entrada"
// 		return c.Status(500).JSON(m)
// 	}
// 	orden := gormdb.Ordenes{}
// 	errdb := db.Find(&orden, "usuario_id = ? AND id = ?", c.Locals("userID"), input.Orden_id)
// 	if errdb.Error != nil {
// 		m["mensaje"] = errdb.Error.Error()
// 		return c.Status(500).JSON(m)
// 	}
// 	return c.JSON(input)
// }

////// ROOT
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
