package shoppingcar

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	pre := "/api" + config.Rest_version + "shoppingcar"

	app.Post(pre+"/shoppingcar", mi.IsRegister, crear)
	app.Get(pre+"/shoppingcar/plan", mi.IsRegister, listarWPlan)
	app.Delete(pre+"/shoppingcar", mi.IsRegister, eliminarall)
	app.Delete(pre+"/shoppingcar/:id", mi.IsRegister, eliminar)

	//TODO no se utilizan aun
	//app.Put(pre+"/carrito", mi.IsRegister, editar)

}
