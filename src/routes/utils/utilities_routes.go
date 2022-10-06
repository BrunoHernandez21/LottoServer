package utils

import (
	"lottomusic/src/config"

	"github.com/gofiber/fiber/v2"
)

//var db *gorm.DB
//sqldb *gorm.DB
func Init_routes(app *fiber.App) {
	//db = sqldb
	pre := "/api" + config.Rest_version + "utils"
	app.Get(pre+"/oclock", oclock)

}
