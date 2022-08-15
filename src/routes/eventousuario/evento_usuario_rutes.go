package eventousuario

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "usuario")

	v1.Post("/evento", mi.IsRegister, crear)
	v1.Get("/evento/:page/:sizepage", mi.IsRegister, historialPage)
	v1.Get("/activos/:page/:sizepage", mi.IsRegister, activosPage)

	v1.Put("/evento", mi.IsRoot, editar)
	v1.Get("/evento", mi.IsRoot, listarTodos)
	v1.Get("/evento/:id", mi.IsRoot, byid)
	v1.Delete("/evento/:id", mi.IsRoot, eliminar)
}
