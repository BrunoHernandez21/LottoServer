package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v3 := app.Group("/api/auth/")
	v2 := app.Group("/api/auth/")
	v1 := app.Group("/api/auth/")
	//v2.Get("logout", logout)

	v1.Post("login", login)
	v1.Post("user", signup)
	v1.Put("forgetpassword", forgetpassword)

	v2.Get("user", infouser)
	v2.Delete("user", deleteuser)
	v2.Put("user", updateuser)
	v2.Post("changepassword", changepassword)
	v2.Get("token", renuevaToken)

	v3.Get("users", users)
	v3.Delete("users/:id", deleteById)
	v3.Get("users/:id", getById)

}
