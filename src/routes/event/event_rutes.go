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
	v1 := app.Group("/api" + config.Rest_version + "event")

	v1.Post("/event", mi.IsRoot, crear)
	v1.Put("/event", mi.IsRoot, editar)
	v1.Delete("/event/:id", mi.IsRoot, eliminar)

	v1.Get("/event/:id", byid)
	v1.Get("/active/:pag/:sizepage", listarActivos)

}
