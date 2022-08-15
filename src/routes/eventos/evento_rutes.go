package evento

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "evento")

	v1.Post("/evento", mi.IsRoot, crear)
	v1.Put("/evento", mi.IsRoot, editar)
	v1.Delete("/evento/:id", mi.IsRoot, eliminar)

	v1.Get("/evento/:id", byid)
	v1.Get("/evento", listarTodos)

	v1.Get("/activos", activo)
	v1.Get("/activos/page/:sizepage", listarActivos)

}
