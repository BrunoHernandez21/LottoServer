package main

import (
	"lottomusic/src/config"
	"lottomusic/src/routes/apuestas"
	"lottomusic/src/routes/auth"
	"lottomusic/src/routes/carrito"
	carritoa "lottomusic/src/routes/carritoapuesta"
	"lottomusic/src/routes/compra"
	"lottomusic/src/routes/juegos"
	"lottomusic/src/routes/planes"
	"lottomusic/src/routes/videos"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
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
	rutasMain(app, db)
	//start Server
	err2 := app.Listen(":25565")
	if err2 != nil {
		panic(err2.Error())
	}

}

func conexionDB() (conexiones *gorm.DB) {
	dns := config.DB.User + ":" + config.DB.Password + config.DB.Soc + config.DB.TableName
	conexion, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func rutasMain(app *fiber.App, db *gorm.DB) {
	auth.Init_routes(app, db)
	planes.Init_routes(app, db)

	carrito.Init_routes(app, db)
	juegos.Init_routes(app, db)
	apuestas.Init_routes(app, db)
	carritoa.Init_routes(app, db)
	compra.Init_routes(app, db)

	videos.Init_routes(app, db)
}

/*
	go mod init
	go mod tidy
	go get -u gorm.io/gorm
*/
