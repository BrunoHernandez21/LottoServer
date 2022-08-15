package carrito

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "carrito")

	v1.Post("/carrito", mi.IsRegister, crear)
	v1.Get("/carrito/plan", mi.IsRegister, listarWPlan)
	v1.Get("/carrito", mi.IsRegister, listar)
	v1.Put("/carrito", mi.IsRegister, editar)
	v1.Delete("/carrito/:id", mi.IsRegister, eliminar)
	v1.Delete("/carritos", mi.IsRegister, eliminarall)

}
