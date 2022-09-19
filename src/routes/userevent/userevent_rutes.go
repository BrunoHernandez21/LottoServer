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
	v1 := app.Group("/api" + config.Rest_version + "user-event")

	//user event
	v1.Post("/event", mi.IsRegister, create_event)
	//history user event
	v1.Get("/active/:page/:sizepage", mi.IsRegister, active_events)
	v1.Get("/history/:page/:sizepage", mi.IsRegister, history_event)
	v1.Get("/wins/:pag/:sizepage", mi.IsRegister, winer_event)

	//root
	v1.Get("byid/:id", mi.IsRoot, event_id)
	v1.Put("/event", mi.IsRoot, userevent_edit)
	v1.Delete("/event/:id", mi.IsRoot, userevent_delete)
	v1.Get("/event/:page/:sizepage", mi.IsRoot, userevent_list_all)

}
