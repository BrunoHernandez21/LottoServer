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
	v1 := app.Group("/api" + config.Rest_version + "video")

	//listVideo
	v1.Get("/video/:id", activoID)
	v1.Get("/videos/:pag/:sizepage", videos_pag)
	v1.Get("/events/:pag/:sizepage", videos_evento_pag)
	// groups
	v1.Get("/groups", listargrupos)
	v1.Get("/groups/:pag/:sizepage/:name", listarGruposName)
	// statistics
	v1.Get("/statistics", get_statistics)
	v1.Get("/statistics/:id", get_st_byID)
	// Root
	v1.Put("/statistics", mi.IsRoot, edit_statistics)
	v1.Delete("/statistics/:id", mi.IsRoot, delete_statistics)
	v1.Post("/video", mi.IsRoot, crear)
	v1.Put("/video", mi.IsRoot, editar)
	v1.Delete("/video/:id", mi.IsRoot, eliminar)
	// v1.Post("/statistics", mi.IsRoot,create_statistics)

}
