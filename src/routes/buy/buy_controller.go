package buy

import (
	"encoding/json"
	"lottomusic/src/models/compuestas"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"lottomusic/src/modules/stripe/impstripe"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//////////////////////////////////////////////////////////////////////////////
///////////////////////// Seccion de Pagos unicos //////////////////////////////
//////////////////////////////////////////////////////////////////////////////
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
		m["mensaje"] = "error al parsear datos de entrada"
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
	data, _ := json.Marshal(&resp)
	db.Raw("CALL orden_pagada(?,?)", orden.Id, string(data)).Scan(&outReason)
	m["resp"] = "Compra realizada con éxito"
	return c.Status(200).JSON(m)
}

//////////////////////////////////////////////////////////////////////////////
///////////////////////// Seccion de Suscripcion //////////////////////////////
//////////////////////////////////////////////////////////////////////////////

//
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

//
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
	if sus.Stripe_suscription != "" {
		m["mensaje"] = "Finalice su suscripcion actual antes de crear una nueva"
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
		err3 := db.Model(&sus).Where("usuario_id = ?", c.Locals("userID")).Updates(&sus)
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
		_, err3 := impstripe.Update_customer(input.Stripe_Payment, sus.Stripe_customer)
		if err3 != nil {
			m["mensaje"] = "Stripe error"
			return c.Status(500).JSON(m)
		}

	}

	stripe_sus, err2 := impstripe.Create_suscription(input.Orden_id, sus.Stripe_customer, *plan.Stripe_price)
	if err2 != nil {
		m["mensaje"] = "Stripe error"
		return c.Status(500).JSON(m)
	}
	sus.Stripe_suscription = stripe_sus.ID
	sus.Usuario_id = 0
	err3 := db.Model(&sus).Where("usuario_id = ?", c.Locals("userID")).Updates(&sus)
	if err3.Error != nil {
		m["mensaje"] = "DB error"
		return c.Status(500).JSON(m)
	}

	m["resp"] = "Compra realizada con éxito"
	return c.Status(200).JSON(m)
}

// Orden de suscripcion
func delete_suscription(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})

	susc := gormdb.Suscripciones{}
	db.Find(&susc, "Usuario_id = ?", c.Locals("userID"))
	if susc.Usuario_id == 0 {
		m["mensaje"] = "No es una suscripcion valida"
		return c.Status(500).JSON(m)
	}
	_, err := impstripe.Delete_suscription(susc.Stripe_suscription)
	if err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	susc.Stripe_suscription = ""
	err3 := db.Model(&susc).Select("Stripe_suscription").Where("usuario_id = ?", c.Locals("userID")).Updates(&susc)
	if err3.Error != nil {
		m["mensaje"] = "DB error"
		return c.Status(500).JSON(m)
	}
	m["resp"] = "Suscripcion finalizada"
	return c.Status(500).JSON(m)
}

func proration_suscription(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})
	input := inputs.SuscripcionCheckout{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "error al parcear datos de entrada"
		return c.Status(500).JSON(m)
	}

	items := gormdb.ItemsOrden{}
	db.Find(&items, "Orden_id = ?", input.Orden_id)
	if items.Plan_id == 0 {
		m["mensaje"] = "No es una orden valida"
		return c.Status(500).JSON(m)
	}
	plan := gormdb.Planes{}
	db.Find(&plan, "id = ?", items.Plan_id)
	if plan.Id == 0 {
		m["mensaje"] = "Error al buscar plan"
		return c.Status(500).JSON(m)
	}
	susc := gormdb.Suscripciones{}
	db.Find(&susc, "Usuario_id = ?", c.Locals("userID"))
	if susc.Usuario_id == 0 {
		m["mensaje"] = "No es una suscripcion valida"
		return c.Status(500).JSON(m)
	}
	item_sub, err := impstripe.Get_suscription(susc.Stripe_suscription)
	if err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if len(item_sub.Items.Data) == 0 {
		m["mensaje"] = "Stripe error"
		return c.Status(500).JSON(m)
	}
	_, err2 := impstripe.Update_suscription_proration(susc.Stripe_suscription, item_sub.Items.Data[0].ID, *plan.Stripe_price, input.Orden_id)
	if err2 != nil {
		m["mensaje"] = err2.Error()
		return c.Status(500).JSON(m)
	}
	susc.Plan_id = plan.Id
	susc.Usuario_id = 0
	err3 := db.Model(&susc).Where("usuario_id = ?", c.Locals("userID")).Updates(&susc)
	if err3.Error != nil {
		m["mensaje"] = "DB error"
		return c.Status(500).JSON(m)
	}

	m["resp"] = "Se cambio la suscripcion"
	return c.Status(500).JSON(m)
}

//
func change_suscription(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})
	input := inputs.SuscripcionCheckout{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "error al parcear datos de entrada"
		return c.Status(500).JSON(m)
	}

	items := gormdb.ItemsOrden{}
	db.Find(&items, "Orden_id = ?", input.Orden_id)
	if items.Plan_id == 0 {
		m["mensaje"] = "No es una orden valida"
		return c.Status(500).JSON(m)
	}
	plan := gormdb.Planes{}
	db.Find(&plan, "id = ?", items.Plan_id)
	if plan.Id == 0 {
		m["mensaje"] = "Error al buscar plan"
		return c.Status(500).JSON(m)
	}
	susc := gormdb.Suscripciones{}
	db.Find(&susc, "Usuario_id = ?", c.Locals("userID"))
	if susc.Usuario_id == 0 {
		m["mensaje"] = "No es una suscripcion valida"
		return c.Status(500).JSON(m)
	}
	item_sub, err := impstripe.Get_suscription(susc.Stripe_suscription)
	if err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if len(item_sub.Items.Data) == 0 {
		m["mensaje"] = "Stripe error"
		return c.Status(500).JSON(m)
	}
	_, err2 := impstripe.Update_suscription_now(susc.Stripe_suscription, item_sub.Items.Data[0].ID, *plan.Stripe_price, input.Orden_id)
	if err2 != nil {
		m["mensaje"] = err2.Error()
		return c.Status(500).JSON(m)
	}
	susc.Plan_id = plan.Id
	susc.Usuario_id = 0
	err3 := db.Model(&susc).Where("usuario_id = ?", c.Locals("userID")).Updates(&susc)
	if err3.Error != nil {
		m["mensaje"] = "DB error"
		return c.Status(500).JSON(m)
	}

	m["resp"] = "Se cambio la suscripcion"
	return c.Status(500).JSON(m)
}

func subscription_change_payment(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})
	input := inputs.SuscripcionCheckout{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "error al parcear datos de entrada"
		return c.Status(500).JSON(m)
	}
	if input.Stripe_Payment == "" {
		m["mensaje"] = "Se requiere un metodo de pago"
		return c.Status(500).JSON(m)
	}
	sus := gormdb.Suscripciones{}
	db.Find(&sus, "usuario_id = ?", c.Locals("userID"))
	if sus.Usuario_id == 0 {
		m["mensaje"] = "Error de usuario"
		return c.Status(500).JSON(m)
	}
	if sus.Usuario_id == 0 || sus.Stripe_customer == "" {
		m["mensaje"] = "Error al buscar suscripcion"
		return c.Status(500).JSON(m)
	}

	_, err2 := impstripe.Atach(sus.Stripe_customer, input.Stripe_Payment)
	if err2 != nil {
		m["mensaje"] = "Stripe error"
		return c.Status(500).JSON(m)
	}
	_, err3 := impstripe.Update_customer(input.Stripe_Payment, sus.Stripe_customer)
	if err3 != nil {
		m["mensaje"] = "Stripe error"
		return c.Status(500).JSON(m)
	}
	impstripe.Detach(sus.Stripe_payment)
	return c.JSON(m)
}

func subscription_get_payment(c *fiber.Ctx) error {
	/// Verificar la respuesta del usuario
	m := make(map[string]interface{})

	sus := gormdb.Suscripciones{}
	db.Find(&sus, "usuario_id = ?", c.Locals("userID"))
	if sus.Usuario_id == 0 {
		m["mensaje"] = "Error de usuario"
		return c.Status(500).JSON(m)
	}
	if sus.Usuario_id == 0 || sus.Stripe_customer == "" {
		m["mensaje"] = "Error al buscar suscripcion"
		return c.Status(500).JSON(m)
	}

	pyment, err := impstripe.Get_paymet_method(sus.Stripe_payment)
	if err != nil {
		return c.Status(500).JSON(m)
	}
	m["last4"] = pyment.Card.Last4
	m["exp_month"] = pyment.Card.ExpMonth
	m["exp_year"] = pyment.Card.ExpYear

	return c.JSON(m)
}

//////////////////////////////////////////////////////////////////////////////
///////////////////////// Seccion de Historial //////////////////////////////
//////////////////////////////////////////////////////////////////////////////

func buy_history_paginated(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	resp := make(map[string]interface{})
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	a := int64(0)
	db.
		Table("pagos_orden").
		Where("usuario_id = ? ", userID).
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
		Where("usuario_id = ? ", c.Locals("userID")).
		Order("fecha_pagado DESC").
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

//////////////////////////////////////////////////////////////////////////////
///////////////////////// Root //////////////////////////////
//////////////////////////////////////////////////////////////////////////////
func eliminar_orden(c *fiber.Ctx) error {
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
