package compute

import (
	"lottomusic/src/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	pre := "/api" + config.Rest_version + "compute"
	// process
	app.Get(pre+"/process/subscriptions", process_subscriptions)
	app.Get(pre+"/process/state-user", process_users)
	app.Get(pre+"/process/statistics", process_statistics)
	app.Get(pre+"/process/winner", process_winner)
	// emit
	app.Get(pre+"/emit/statistics", emit_statistics)
	app.Get(pre+"/emit/winner", emit_winner)
	// Webhook
	app.Post(pre+"/webhook/stripe", stripe_webhook)

}
