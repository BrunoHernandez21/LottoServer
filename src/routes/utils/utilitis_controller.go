package utils

import (
	"lottomusic/src/models/gormdb"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func oclock(c *fiber.Ctx) error {
	m := make(map[string]string)
	m["time"] = time.Now().Local().String()[0:25]
	return c.JSON(m)
}

func ganador(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	resp := make(map[string]interface{})

	a := int64(0)
	db.Table("ganador").Where("Id_usuario = ?", userID).Count(&a)
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

	ganador := []gormdb.Ganador{}
	errdb := db.Table("ganador").Offset(int(init)).Limit(int(sizepage)).Find(&ganador, "Id_usuario = ?", userID)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}

	resp["ganador"] = ganador

	return c.JSON(resp)
}

func cartera(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	var user = c.Locals("userID")
	userID, ok := user.(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}

	cartera := gormdb.Carteras{}
	errdb := db.Find(&cartera, "Id_usuario = ?", userID)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if cartera.Id == 0 {
		cartera.Id_usuario = userID
		errdb = db.Create(&cartera)
		if errdb.Error != nil {
			m["mensaje"] = errdb.Error.Error()
			return c.Status(500).JSON(m)
		}
	}
	errdb = db.Find(&cartera, "Id_usuario = ?", userID)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["cartera"] = cartera
	return c.JSON(m)
}
