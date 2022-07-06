package utils

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init_routes(app *fiber.App, sqldb *gorm.DB) {

	v1 := app.Group("/api/utils")
	v1.Get("/oclock", oclock)
}
