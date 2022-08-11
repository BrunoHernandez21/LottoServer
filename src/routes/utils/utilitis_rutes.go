package utils

import (
	"lottomusic/src/config"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "utils")
	v1.Get("/oclock", oclock)
	v1.Get("/wins/:pag/:sizepage", isRegister, ganador)
	v1.Get("/cartera", isRegister, cartera)
}

func isRegister(c *fiber.Ctx) error {
	m := make(map[string]string)
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m["mensaje"] = "Token invalido"
		return c.Status(500).JSON(m)
	}
	user := gormdb.Usuarios{}
	errdb := db.Find(&user, "id = ?", credentials.ID)
	if errdb.Error != nil {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	if user.Id == 0 {
		m["mensaje"] = "Usuario no registado"
		return c.Status(500).JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}
