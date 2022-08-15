package utils

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "utils")
	v1.Get("/oclock", oclock)
	v1.Get("/wins/:pag/:sizepage", mi.IsRegister, ganador)
	v1.Get("/cartera", mi.IsRegister, cartera)
}
