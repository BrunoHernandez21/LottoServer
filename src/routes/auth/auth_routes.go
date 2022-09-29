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

	v1 := app.Group("/api" + config.Rest_version + "auth/")
	//v2.Get("logout", logout)

	v1.Post("login", login)
	v1.Post("user", signup)
	v1.Put("forgetpassword", forgetpassword)
	v1.Get("token", mi.IsRegister, renewToken)

	v1.Get("users", mi.IsRoot, users)
	v1.Delete("users/:id", mi.IsRoot, deleteById)
	v1.Get("users/:id", mi.IsRoot, getById)

}
