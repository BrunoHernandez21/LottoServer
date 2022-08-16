package videos

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "videos")

	v1.Get("/eventos/:page/:sizepage", videos_evento_pag)

	v1.Get("/video/:id", activoID)
	v1.Get("/videos/:page/:sizepage", videos_pag)
	v1.Post("/videos", mi.IsRoot, crear)
	v1.Put("/videos", mi.IsRoot, editar)
	v1.Delete("/videos/:id", mi.IsRoot, eliminar)

	v1.Get("/grupos", listargrupos)
	v1.Get("/grupos/:page/:sizepage/:name", listarGruposName)

	v1.Get("/estadisticas", get_statistics)
	v1.Get("/estadisticas/:id", get_st_byID)
	v1.Post("/estadisticas", mi.IsRoot, create_statistics)
	v1.Put("/estadisticas", mi.IsRoot, edit_statistics)
	v1.Delete("/estadisticas/:id", mi.IsRoot, delete_statistics)
}
