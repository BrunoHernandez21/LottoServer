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
	pre := "/api" + config.Rest_version + "video"

	//listVideo
	app.Get(pre+"/video/:id", activoID)
	app.Get(pre+"/videos/:pag/:sizepage", videos_pag)
	app.Get(pre+"/events/:pag/:sizepage", videos_evento_pag)
	// groups
	app.Get(pre+"/groups", listargrupos)
	app.Get(pre+"/groups/:pag/:sizepage/:name", listarGruposName)
	// statistics
	app.Get(pre+"/statistics", get_statistics)
	app.Get(pre+"/statistics/:id", get_st_byID)
	// Root
	app.Put(pre+"/statistics", mi.IsRoot, edit_statistics)
	app.Delete(pre+"/statistics/:id", mi.IsRoot, delete_statistics)
	app.Post(pre+"/video", mi.IsRoot, crear)
	app.Put(pre+"/video", mi.IsRoot, editar)
	app.Delete(pre+"/video/:id", mi.IsRoot, eliminar)
	// app.Post("/statistics", mi.IsRoot,create_statistics)

}
