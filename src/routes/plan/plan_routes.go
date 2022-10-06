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
	pre := "/api" + config.Rest_version + "plan"

	app.Get(pre+"/plans/single-payment", single_payment)
	app.Get(pre+"/plans/suscripcion", list_subscriptions)
	app.Get(pre+"/byname/:name", byname)
	app.Get(pre+"/byid/:id", byid)

	app.Post(pre+"/plan", mi.IsRoot, create)
	app.Delete(pre+"/plan/:id", mi.IsRoot, delete)
	app.Put(pre+"/plan", mi.IsRoot, edit)

}
