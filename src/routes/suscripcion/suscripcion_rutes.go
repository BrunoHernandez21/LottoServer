package suscripcion

import (
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/suscripcion")
	v1.Get("/suscripcion", isRegister, listar)
	v1.Get("/suscripcion/activo", isRegister, listaractivos)

}

func isRegister(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m := make(map[string]string)
		m["mensjae"] = "Token invalido"
		return c.JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}
