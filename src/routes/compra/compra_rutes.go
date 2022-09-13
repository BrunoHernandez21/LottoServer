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

	// manda la orden
	v1.Post("/checkout", mi.IsRegister, checkout)
	// Punto de acceso para stripe
	v1.Post("/verifica", verifica)
	//Historial de compras y ordenes
	v1.Get("/compra", mi.IsRegister, listar)
	v1.Get("/compra/:pag/:sizepage", mi.IsRegister, listarpaginado)
	v1.Get("/ordens/status", mi.IsRegister, listarOrdenes)

	//Tarjetas y su administracion
	v1.Post("/payment/method", mi.IsRegister, createTarjeta)
	v1.Put("/payment/method", mi.IsRegister, editTarjeta)
	v1.Delete("/payment/method/:id", mi.IsRegister, deleteTarjeta)
	v1.Get("/payment/method", mi.IsRegister, listarTarjeta)

	//ROOT
	v1.Delete("/compra", mi.IsRoot, eliminar)
}
