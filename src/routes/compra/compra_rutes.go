package compra

import (
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb

	v1 := app.Group("/api/compra")

	v1.Post("/compra", isRegister, crear)
	v1.Delete("/compra", isRegister, eliminar)
	v1.Get("/compra", isRegister, listar)
	v1.Post("/checkout", isRegister, checkout)
}

func isRegister(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m := make(map[string]string)
		m["mensaje"] = "Token invalido"
		return c.JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}
