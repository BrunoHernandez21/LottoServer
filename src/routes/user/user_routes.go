package user

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api" + config.Rest_version + "user")

	//user
	v1.Get("user", mi.IsRegister, infouser)
	v1.Delete("user", mi.IsRegister, deleteuser)
	v1.Put("user", mi.IsRegister, updateuser)
	v1.Put("changepassword", mi.IsRegister, changepassword)

	// direction
	v1.Post("direction", mi.IsRegister, createDireccion)
	v1.Get("direction", mi.IsRegister, getDireccion)
	v1.Put("direction", mi.IsRegister, updateDireccion)
	v1.Delete("direction/:id", mi.IsRegister, deleteDireccion)

	//user-state
	v1.Get("/purse", mi.IsRegister, cartera)
	v1.Get("/properties", mi.IsRegister, propiedades)
	v1.Get("/subscription", mi.IsRegister, suscribcion)

	// payment-method
	v1.Post("/payment/method", mi.IsRegister, createTarjeta)
	v1.Put("/payment/method", mi.IsRegister, editTarjeta)
	v1.Delete("/payment/method/:id", mi.IsRegister, deleteTarjeta)
	v1.Get("/payment/method", mi.IsRegister, listarTarjeta)

}
