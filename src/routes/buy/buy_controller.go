package buy

import (
	"encoding/json"
	"fmt"
	"lottomusic/src/models/compuestas"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"lottomusic/src/modules/stripe/impstripe"
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
	db.Raw("CALL genera_orden(?)", c.Locals("userID")).Scan(&orden)
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
	a, err := impstripe.Create_payment_intent(&orden)
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
	resp, errstr := impstripe.Pay_payment_intent(&orden, input.Stripe_Payment)
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
	m := make(map[string]interface{})
	input := inputs.SuscripcionOrden{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "error al parcear datos de entrada"
		return c.Status(500).JSON(m)
	}
	orden := gormdb.Ordenes{}
	db.Raw("CALL orden_subscripcion( ? , ? )", c.Locals("userID"), input.Plan_id).Scan(&orden)
	if orden.Id == 0 {
		m["mensaje"] = "No es una suscripcion valida"
		return c.Status(500).JSON(m)
	}

	items_orden := []gormdb.ItemsOrden{}
	db.Find(&items_orden, "Orden_id = ?", orden.Id)
	m["orden"] = orden
	m["items_orden"] = items_orden
	return c.Status(500).JSON(m)
}

func subscription_checkout(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})
	input := inputs.SuscripcionCheckout{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "error al parcear datos de entrada"
		return c.Status(500).JSON(m)
	}
	itms_ord := gormdb.ItemsOrden{}
	db.Find(&itms_ord, "Orden_id = ?", input.Orden_id)
	plan := gormdb.Planes{}
	db.Find(&plan, "id = ?", itms_ord.Plan_id)
	if plan.Id == 0 || plan.Stripe_price == nil || !plan.Suscribcion {
		m["mensaje"] = "este plan no cumple con los requisitos"
		return c.Status(500).JSON(m)
	}
	sus := gormdb.Suscripciones{}
	db.Find(&sus, "usuario_id = ?", c.Locals("userID"))
	if sus.Usuario_id == 0 {
		m["mensaje"] = "Error de usuario"
		return c.Status(500).JSON(m)
	}
	if sus.Stripe_customer == "" {
		cus, err2 := impstripe.Create_customer(input.Stripe_Payment, sus.Usuario_id)
		if err2 != nil {
			m["mensaje"] = "Stripe error"
			return c.Status(500).JSON(m)
		}
		sus.Stripe_payment = input.Stripe_Payment
		sus.Stripe_customer = cus.ID
		sus.Usuario_id = 0
		fmt.Println("pre incert")
		err3 := db.Model(&sus).Where("usuario_id = ?", c.Locals("userID")).Updates(sus)
		if err3.Error != nil {
			m["mensaje"] = "DB error"
			return c.Status(500).JSON(m)
		}
	} else {
		impstripe.Detach(sus.Stripe_payment)
		_, err2 := impstripe.Atach(sus.Stripe_customer, input.Stripe_Payment)
		if err2 != nil {
			m["mensaje"] = "Stripe error"
			return c.Status(500).JSON(m)
		}
	}

	fmt.Println("Salgo del if")
	stripe_sus, err2 := impstripe.Create_suscription(input.Orden_id, sus.Stripe_customer, *plan.Stripe_price)
	if err2 != nil {
		m["mensaje"] = "Stripe error"
		return c.Status(500).JSON(m)
	}
	fmt.Println("if")
	sus.Stripe_suscription = stripe_sus.ID
	sus.Usuario_id = 0
	err3 := db.Model(&sus).Where("usuario_id = ?", c.Locals("userID")).Updates(sus)
	if err3.Error != nil {
		m["mensaje"] = "DB error"
		return c.Status(500).JSON(m)
	}

	m["resp"] = "Compra realizada con éxito"
	return c.Status(200).JSON(m)
}

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
