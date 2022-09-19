package utils

import (
	"lottomusic/src/config"

	"github.com/gofiber/fiber/v2"
)

//var db *gorm.DB
//sqldb *gorm.DB
func Init_routes(app *fiber.App) {
	//db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "utils")
	v1.Get("/oclock", oclock)

}
