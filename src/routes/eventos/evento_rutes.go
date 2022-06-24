package evento

import (
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/evento")

	v1.Post("/evento", isRoot, crear)
	v1.Put("/evento", isRoot, editar)
	v1.Get("/evento/:id", isRoot, byid)
	v1.Get("/evento/all", isRoot, listarTodos)
	v1.Delete("/evento/:id", isRoot, eliminar)

	v1.Post("/activos", isRegister, activo)
	v1.Get("/activos/page/:page", isRegister, listarActivos)

}

func isRegister(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m := make(map[string]string)
		m["mensjae"] = "Token invalido"
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
		m["mensjae"] = "Token invalido"
		return c.JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}
