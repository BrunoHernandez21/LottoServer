package plan

import (
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/plan")

	v1.Get("/planes", isRegister, lista)
	v1.Get("/byname/:name", isRegister, byname)
	v1.Get("/byid/:id", isRegister, byid)

	v1.Post("/plan", isRoot, create)
	v1.Delete("/plan/:id", isRoot, delete)
	v1.Put("/plan", isRoot, edit)

}

func isRegister(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m := make(map[string]string)
		m["mensaje"] = "Token invalido"
		return c.JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}

func isRoot(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m := make(map[string]string)
		m["mensaje"] = "Token invalido"
		return c.JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}
