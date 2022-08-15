package main

import (
	"encoding/json"
	"io/ioutil"
	"lottomusic/src/config"
	"lottomusic/src/modules/midelware"
	"lottomusic/src/routes/auth"
	"lottomusic/src/routes/carrito"
	"lottomusic/src/routes/compra"
	evento "lottomusic/src/routes/eventos"
	"lottomusic/src/routes/eventousuario"
	"lottomusic/src/routes/plan"
	"lottomusic/src/routes/suscripcion"
	"lottomusic/src/routes/utils"
	"lottomusic/src/routes/videos"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//LoadFiles
	loadInitialConfig()
	//mainConfig
	app := fiber.New(fiber.Config{
		AppName: "Loto Music",
		Prefork: false,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Content-Length, Content-Type, Authorization, Accept, Origin",
		ExposeHeaders:    "Content-Length, Content-Type, Authorization, Accept, Origin",
		AllowCredentials: true,
		MaxAge:           1800,
	}))

	//instance DB
	db := conexionDB()
	//instance Ruts
	app.Static("/pagina/web", "./home.html")
	rutasMain(app, db)
	//start Server
	err2 := app.Listen(":" + config.Rest_Port)
	if err2 != nil {
		panic(err2.Error())
	}

}

func conexionDB() (conexiones *gorm.DB) {
	dns := config.DB.User + ":" + config.DB.Password + config.DB.Soc + config.DB.DatabaseName + "?charset=utf8&parseTime=True"
	conexion, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func rutasMain(app *fiber.App, db *gorm.DB) {
	midelware.Init_state(db)
	auth.Init_routes(app, db)
	plan.Init_routes(app, db)
	carrito.Init_routes(app, db)
	compra.Init_routes(app, db)
	eventousuario.Init_routes(app, db)
	evento.Init_routes(app, db)
	suscripcion.Init_routes(app, db)
	videos.Init_routes(app, db)
	utils.Init_routes(app, db)

}

func loadInitialConfig() {
	db_file, err := ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		panic(err.Error())
	}
	data := ConfigMain{}
	errjs := json.Unmarshal(db_file, &data)
	if errjs != nil {
		panic(err.Error())
	}
	config.DB = data.MainDB
	config.Mail = data.MainMail
	config.Rest_Port = data.Rest_Port
	config.Soc_Port = data.Soc_Port
}

type ConfigMain struct {
	MainDB    config.ConfigDB    `json:"db"`
	MainMail  config.ConfigEmail `json:"mail_config"`
	Rest_Port string             `json:"port_rest"`
	Soc_Port  string             `json:"port_soc"`
}

/*
	go mod init
	go mod tidy
	go get -u gorm.io/gorm
*/
