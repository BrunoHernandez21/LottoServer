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
	pre := "/api" + config.Rest_version + "user"

	//user
	app.Get(pre+"user", mi.IsRegister, infouser)
	app.Delete(pre+"user", mi.IsRegister, deleteuser)
	app.Put(pre+"user", mi.IsRegister, updateuser)
	app.Put(pre+"changepassword", mi.IsRegister, changepassword)

	// direction
	app.Post(pre+"direction", mi.IsRegister, createDireccion)
	app.Get(pre+"direction", mi.IsRegister, getDireccion)
	app.Put(pre+"direction", mi.IsRegister, updateDireccion)
	app.Delete(pre+"direction/:id", mi.IsRegister, deleteDireccion)

	//user-state
	app.Get(pre+"/purse", mi.IsRegister, cartera)
	app.Get(pre+"/properties", mi.IsRegister, propiedades)
	app.Get(pre+"/subscription", mi.IsRegister, suscripcion)

	// payment-method
	app.Post(pre+"/payment/method", mi.IsRegister, createTarjeta)
	app.Put(pre+"/payment/method", mi.IsRegister, editTarjeta)
	app.Delete(pre+"/payment/method/:id", mi.IsRegister, deleteTarjeta)
	app.Get(pre+"/payment/method", mi.IsRegister, listarTarjeta)

}
