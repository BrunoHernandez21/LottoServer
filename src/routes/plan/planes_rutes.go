package plan

import (
	"lottomusic/src/models/gormdb"
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
	m := make(map[string]string)
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m["mensaje"] = "Token invalido"
		return c.Status(500).JSON(m)
	}
	user := gormdb.Usuarios{}
	errdb := db.Find(&user, "id = ?", credentials.ID)
	if errdb.Error != nil {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	if user.Id == 0 {
		m["mensaje"] = "Usuario no registado"
		return c.Status(500).JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}

func isRoot(c *fiber.Ctx) error {
	m := make(map[string]string)
	headers := c.GetReqHeaders()
	_, credentials, err := jwts.ValidateToken(headers["Authorization"])
	if err != nil {
		m["mensaje"] = "Token invalido"
		return c.Status(500).JSON(m)
	}
	user_rol := gormdb.Usuarios_roles{}
	errdb := db.Find(&user_rol, "User_id = ?", credentials.ID)
	if errdb.Error != nil {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	if user_rol.User_id == 0 {
		m["mensaje"] = "Usuario no registado"
		return c.Status(500).JSON(m)
	}
	if user_rol.Role_id != 2 {
		m["mensaje"] = "Unauthorized"
		return c.Status(401).JSON(m)
	}
	c.Locals("userID", credentials.ID)
	return c.Next()
}
