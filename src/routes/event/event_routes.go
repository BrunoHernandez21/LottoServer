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
	pre := "/api" + config.Rest_version + "event"

	app.Post(pre+"/event", mi.IsRoot, crear)
	app.Put(pre+"/event", mi.IsRoot, editar)
	app.Delete(pre+"/event/:id", mi.IsRoot, eliminar)

	app.Get(pre+"/event/:id", byid)
	app.Get(pre+"/active/:pag/:sizepage", listarActivos)

}
