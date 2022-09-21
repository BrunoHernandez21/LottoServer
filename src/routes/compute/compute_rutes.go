package compute

import (
	"lottomusic/src/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "compute")
	// process
	v1.Get("/process/subscriptions", process_subscriptions)
	v1.Get("/process/state-user", process_users)
	v1.Get("/process/statistics", process_statistics)
	v1.Get("/process/winner", process_winner)
	// emit
	v1.Get("/emit/statistics", emit_statistics)
	v1.Get("/emit/winner", emit_winner)
	// Webhook
	v1.Post("/webhook/stripe", stripe_webhook)

}
