package compra

import (
	"lottomusic/src/models/compra"
	"lottomusic/src/models/gormdb"
	"math"
	"strconv"
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
		m["mensaje"] = err.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}
func listar(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := []gormdb.Compra{}
	errdb := db.Find(&input, "Usuario_id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}

func listarpaginado(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	resp := make(map[string]interface{})

	a := int64(0)
	db.Table("Compra").Where("Usuario_id = ?", userID).Count(&a)
	pag, err := strconv.ParseUint(c.Params("pag"), 0, 32)
	sizepage, err2 := strconv.ParseUint(c.Params("sizepage"), 0, 32)
	if err != nil || err2 != nil {
		resp["mensaje"] = err.Error()
		return c.Status(500).JSON(resp)
	}
	pags := math.Round(float64(a) / float64(sizepage))
	if pags < 1 && a > 0 {
		pags = 1
	}
	resp["pags"] = pags
	resp["pag"] = &pag
	resp["sizePage"] = &sizepage
	resp["totals"] = &a
	init := (pag - 1) * sizepage

	compra := []gormdb.Compra{}
	errdb := db.Table("Compra").Offset(int(init)).Limit(int(sizepage)).Find(&compra, "Usuario_id = ?", userID)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["compras"] = compra

	return c.JSON(resp)

}

func checkout(c *fiber.Ctx) error {
	m := make(map[string]interface{})

	input := compra.Get_Checkout{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
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
		m["mensjae"] = errdb.Error.Error()
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
			m["mensaje"] = errdb.Error.Error()
			return c.Status(500).JSON(m)
		}
	}

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("POST")
	a.ContentType("application/json")
	a.JSON(input)
	//req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNjU2NzE1NjcxfQ.4bdSgvBkYL_b8N15nUsQz2r5F1ORCHPrsdMcYJVnsBQaqACxgKk0Fa9YWbJvOQ_MTwKX4MTZgLmUzRyBKyE9Zw")
	req.SetRequestURI("http://187.213.79.204:25565/api/compra/compra")
	if err := a.Parse(); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	code, body, errs := a.Bytes()
	m["mensaje"] = string(body)
	m["errors"] = errs
	return c.Status(code).JSON(m)
}
