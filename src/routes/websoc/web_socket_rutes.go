package socrute

import (
	"lottomusic/src/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	v1 := app.Group("/api" + config.Rest_version + "soc")

	v1.Post("/emitir", emitir)

}
