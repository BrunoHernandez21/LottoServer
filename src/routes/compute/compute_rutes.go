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
	v1.Get("/process/subscriptions", subscriptions)
	v1.Get("/process/state-user", stateusers)
	v1.Get("/process/statistics", statistics)
	v1.Get("/process/winner", winner)
	// stripe
	v1.Post("/webhook/stripe", statistics)
	// emit
	v1.Get("/emit/statistics", emit)
	v1.Get("/emit/winner", winner)

}
