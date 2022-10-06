package auth

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	pre := "/api" + config.Rest_version + "auth/"

	app.Post(pre+"login", login)
	app.Post(pre+"user", signup)
	app.Put(pre+"forgetpassword", forgetpassword)
	app.Get(pre+"token", mi.IsRegister, renewToken)

	app.Get(pre+"users", mi.IsRoot, users)
	app.Delete(pre+"users/:id", mi.IsRoot, deleteById)
	app.Get(pre+"users/:id", mi.IsRoot, getById)

	app.Get("/hola/guapo", func(c *fiber.Ctx) error {
		return c.SendString("Secure connection")
	})
}
