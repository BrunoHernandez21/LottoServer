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

	//
	v1.Get("/subscriptions", subscriptions)
	v1.Get("/stateusers", stateusers)
	v1.Post("/statistics", statistics)
	v1.Get("/winner", winner)
	v1.Get("/emit", emit)

}
