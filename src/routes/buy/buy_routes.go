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
	v1.Post("/orders/order", mi.IsRegister, create_order)
	v1.Post("/orders/payment-intent", mi.IsRegister, create_payment_intent)
	v1.Post("/orders/checkout", mi.IsRegister, checkout)
	// suscripciónes
	v1.Delete("/orders/subscription", mi.IsRegister, delete_suscription)
	v1.Put("/orders/subscription/proration", mi.IsRegister, proration_suscription)
	v1.Post("/orders/subscription/orden", mi.IsRegister, subscription_orden)
	v1.Post("/orders/subscription/checkout", mi.IsRegister, subscription_checkout)
	// history
	v1.Get("/orders/waiting", mi.IsRegister, list_orders)
	v1.Get("/orders/rejected", mi.IsRegister, list_orders_errors)
	v1.Get("/history/:pag/:sizepage", mi.IsRegister, buy_history_paginated)
	// ROOT
	v1.Delete("/orders/:id", mi.IsRoot, eliminar)

	// v1.Post("/orders/rentry", mi.IsRegister, buy_retry)
	// v1.Post("/orders/cancel", mi.IsRegister, buy_cancel)
}
