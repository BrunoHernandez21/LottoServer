package plan

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "plan")

	v1.Get("/plans/single-payment", single_payment)
	v1.Get("/plans/suscripcion", list_subscriptions)
	v1.Get("/byname/:name", byname)
	v1.Get("/byid/:id", byid)

	v1.Post("/plan", mi.IsRoot, create)
	v1.Delete("/plan/:id", mi.IsRoot, delete)
	v1.Put("/plan", mi.IsRoot, edit)

}
