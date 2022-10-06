package userevent

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	pre := "/api" + config.Rest_version + "user-event"

	//user event
	app.Post(pre+"/event", mi.IsRegister, create_event)
	//history user event
	app.Get(pre+"/active/:pag/:sizepage", mi.IsRegister, active_events)
	app.Get(pre+"/history/:pag/:sizepage", mi.IsRegister, history_event)
	app.Get(pre+"/wins/:pag/:sizepage", mi.IsRegister, winer_event)

	//root
	app.Get(pre+"byid/:id", mi.IsRoot, event_id)
	app.Put(pre+"/event", mi.IsRoot, userevent_edit)
	app.Delete(pre+"/event/:id", mi.IsRoot, userevent_delete)
	// app.Get(pre+"/event/:pag/:sizepage", mi.IsRoot, userevent_list_all)

}
