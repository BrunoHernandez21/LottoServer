package compra

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb

	v1 := app.Group("/api" + config.Rest_version + "compra")
	//TODO esto deberia ser root

	v1.Delete("/compra", mi.IsRoot, eliminar)

	v1.Post("/verifica", mi.IsRegister, verifica)
	v1.Get("/compra", mi.IsRegister, listar)
	v1.Get("/compra/:pag/:sizepage", mi.IsRegister, listarpaginado)
	v1.Post("/checkout", mi.IsRegister, checkout)

	v1.Post("/payment/method", mi.IsRegister, createTarjeta)
	v1.Put("/payment/method", mi.IsRegister, editTarjeta)
	v1.Delete("/payment/method/:id", mi.IsRegister, deleteTarjeta)
	v1.Get("/payment/method", mi.IsRegister, listarTarjeta)

	v1.Get("/ordens/status", mi.IsRegister, listarOrdenes)
}
