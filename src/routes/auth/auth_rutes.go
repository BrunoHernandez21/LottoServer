package auth

import (
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/auth/")
	//v2.Get("logout", logout)

	v1.Post("login", login)
	v1.Post("user", signup)
	v1.Put("forgetpassword", forgetpassword)

	v1.Get("user", middleware1, infouser)
	v1.Delete("user", middleware1, deleteuser)
	v1.Put("user", middleware1, updateuser)
	v1.Put("changepassword", middleware1, changepassword)
	v1.Get("token", middleware1, renuevaToken)

	v1.Get("users", middleware1, users)
	v1.Delete("users/:id", middleware1, deleteById)
	v1.Get("users/:id", middleware1, getById)

}

func middleware1(c *fiber.Ctx) error {
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
