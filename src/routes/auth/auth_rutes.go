package auth

import (
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/auth/")
	//v2.Get("logout", logout)

	v1.Post("login", login)
	v1.Post("user", signup)
	v1.Put("forgetpassword", forgetpassword)

	v1.Get("user", mi.IsRegister, infouser)
	v1.Delete("user", mi.IsRegister, deleteuser)
	v1.Put("user", mi.IsRegister, updateuser)
	v1.Put("changepassword", mi.IsRegister, changepassword)
	v1.Get("token", mi.IsRegister, renuevaToken)

	v1.Get("direccion", mi.IsRegister, getDireccion)
	v1.Delete("direccion/:id", mi.IsRegister, deleteDireccion)
	v1.Put("direccion", mi.IsRegister, updateDireccion)
	v1.Post("direccion", mi.IsRegister, createDireccion)

	v1.Get("users", mi.IsRoot, users)
	v1.Delete("users/:id", mi.IsRoot, deleteById)
	v1.Get("users/:id", mi.IsRoot, getById)

}
