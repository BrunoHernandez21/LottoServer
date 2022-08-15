package suscripcion

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "suscripcion")
	v1.Get("/suscripcion", mi.IsRegister, listar)
	v1.Get("/suscripcion/activo", mi.IsRegister, listaractivos)
	v1.Get("/suscripcion/all", mi.IsRegister, listarall)
}
