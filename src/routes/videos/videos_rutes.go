package videos

import (
	"lottomusic/src/models/gormdb"
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/videos")

	v1.Get("/activo/:page/:sizepage", pagelistar)
	v1.Get("/activo", listaractivos)
	v1.Get("/grupos", listargrupos)
	v1.Get("/videos/:id", activoID)

	v1.Post("/videos", isRoot, crear)
	v1.Get("/videos", isRoot, listar)
	v1.Put("/videos", isRoot, editar)
	v1.Delete("/videos/:id", isRoot, eliminar)
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
