package compra

import (
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := inputs.Checkout{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	carritos := []gormdb.Carrito{}
	errdb := db.Where("id IN ?", input.IDs).Find(&carritos)
	if errdb.Error != nil {
		m["mensjae"] = "Internal Error"
		return c.Status(500).JSON(m)
	}
	for _, carrito := range carritos {
		//retirar del carrtio
		activo := false
		carrito.Activo = &activo
		errdb := db.Save(&carrito)
		if errdb.Error != nil {
			m["mensaje"] = "Error interno"
			return c.Status(500).JSON(m)
		}
		//update cartera
		cartera := gormdb.Carteras{
			Id: carrito.Usuario_id,
		}
		errdb = db.Find(&cartera)
		if errdb.Error != nil {
			m["mensaje"] = "Error interno"
			return c.Status(500).JSON(m)
		}
		plan := gormdb.Planes{
			Id: carrito.Plan_id,
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
			Fecha_pagado: &tiempo,
			Carrito_id:   carrito.Id,
			Usuario_id:   carrito.Usuario_id,
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
	db.Table("compra").Where("Usuario_id = ?", userID).Count(&a)
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
	errdb := db.Table("compra").Offset(int(init)).Limit(int(sizepage)).Find(&compra, "Usuario_id = ?", userID)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["compras"] = compra

	return c.JSON(resp)

}

func checkout(c *fiber.Ctx) error {
	m := make(map[string]interface{})

	input := inputs.Checkout{}
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

	carritos := []gormdb.Carrito{}
	errdb := db.Where("id IN ?", input.IDs).Find(&carritos)
	if errdb.Error != nil {
		m["mensjae"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	for _, carrito := range carritos {
		if !*carrito.Activo {
			m["mensaje"] = "Esta orden ha expirado"
			m["OrderID"] = carrito.Id
			return c.Status(500).JSON(m)
		}
		if carrito.Usuario_id != userID {
			m["mensaje"] = "Esta orden no te pertenece"
			m["OrderID"] = carrito.Id
			return c.Status(500).JSON(m)
		}
		errdb := db.Save(&carrito)
		if errdb.Error != nil {
			m["mensaje"] = errdb.Error.Error()
			return c.Status(500).JSON(m)
		}
	}
	/*
		a := fiber.AcquireAgent()
		req := a.Request()
		req.Header.SetMethod("POST")
		a.ContentType("application/json")
		a.JSON(input)
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6Mn0.dzkt3xFC_V-DSxufwf9VsBlMWM2JEeWkLmTLS3Tc6TeyhnodMX4OOIeGcd1SOAj320EUbA7bvIzJ7btAdhU1oA")
		req.SetRequestURI("http://187.213.77.165:25565/api/compra/compra")
		if err := a.Parse(); err != nil {
			m["mensaje"] = err.Error()
			return c.Status(500).JSON(m)
		}
		code, body, errs := a.Bytes()*/
	m["mensaje"] = devfunc(input.IDs)
	m["errors"] = nil
	if m["mensaje"] != "correcto" {
		return c.Status(500).JSON(m)
	}
	return c.Status(200).JSON(m)
}

/// pre stripe
func devfunc(input []uint32) string {

	carritos := []gormdb.Carrito{}
	errdb := db.Where("id IN ?", input).Find(&carritos)
	if errdb.Error != nil {

		return "error"
	}
	for _, carrito := range carritos {
		//retirar del carrtio
		activo := false
		carrito.Activo = &activo
		errdb := db.Save(&carrito)
		if errdb.Error != nil {

			return "error"
		}
		//update cartera
		cartera := gormdb.Carteras{
			Id: carrito.Usuario_id,
		}
		errdb = db.Find(&cartera)
		if errdb.Error != nil {
			return "error"
		}
		plan := gormdb.Planes{
			Id: carrito.Plan_id,
		}
		errdb = db.Find(&plan)
		if errdb.Error != nil {
			return "error"
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
			Fecha_pagado: &tiempo,
			Usuario_id:   carrito.Usuario_id,
			Carrito_id:   carrito.Id,
		}
		errdb = db.Create(&compra)
		if errdb.Error != nil {
			return "error"
		}
	}
	return "correcto"
}
