package apuesta

import (
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/apuesta")

	v1.Post("/apuesta", isRegister, crear)
	v1.Put("/apuesta", isRegister, editar)
	v1.Get("/apuesta", isRegister, listarTodos)

	v1.Get("/activos", isRegister, activo)
	v1.Get("/activos/page/:page", isRegister, listarActivos)

	v1.Get("/apuesta/:id", isRegister, byid)
	v1.Delete("/apuesta/:id", isRoot, eliminar)

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
