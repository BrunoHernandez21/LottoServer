package apuestas

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v3 := app.Group("/api/apuesta/")
	v2 := app.Group("/api/apuesta/")

	v2.Get("add", logout)
	v2.Get("active", login)
	v2.Get("list", renuevaToken)
	v3.Get("activepage/:page", renuevaToken)

}
