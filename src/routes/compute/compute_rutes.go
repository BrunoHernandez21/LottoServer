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

	//comprar
	v1.Post("/statistics", statistics)
	// Punto de acceso para stripe
	v1.Post("/winner", winner)

}
