package shoppingcar

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "shoppingcar")

	v1.Post("/shoppingcar", mi.IsRegister, crear)
	v1.Get("/shoppingcar/plan", mi.IsRegister, listarWPlan)
	v1.Delete("/shoppingcar", mi.IsRegister, eliminarall)
	v1.Delete("/shoppingcar/:id", mi.IsRegister, eliminar)

	//TODO no se utilizan aun
	//v1.Put("/carrito", mi.IsRegister, editar)

}
