package buy

import (
	"lottomusic/src/config"
	mi "lottomusic/src/modules/midelware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb

	v1 := app.Group("/api" + config.Rest_version + "buy")

	// manda la orden
	v1.Post("/orders/checkout", mi.IsRegister, checkout)
	v1.Post("/orders/rentry", mi.IsRegister, buy_retry)
	v1.Post("/orders/cancel", mi.IsRegister, buy_cancel)
	//history
	v1.Get("/orders/waiting", mi.IsRegister, list_orders)
	v1.Get("/orders/rejected", mi.IsRegister, list_orders_errors)
	v1.Get("/history/:pag/:sizepage", mi.IsRegister, buy_history_paginated)

	//ROOT
	v1.Delete("/orders/:id", mi.IsRoot, eliminar)
}
