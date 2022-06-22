package auth

import (
	"crypto/sha1"
	"encoding/hex"
	gormdb "lottomusic/src/models/gormDB"
	"lottomusic/src/models/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func login(c *fiber.Ctx) error {
	//catch midelware
	input := services.Auth_Get_Login{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if (input.Username == nil) || (input.Password == nil) {
		m := make(map[string]string)
		m["mensjae"] = "informacion insuficiente"
		return c.JSON(m)
	}

	//db midelware
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "email = ?", input.Username)
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}
	//password midelware
	h := sha1.New()
	h.Write([]byte(*input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = &i
	//validator midelware
	if a.Password != *input.Password {
		m := make(map[string]string)
		m["mensjae"] = "Contraseña invalida"
		return c.JSON(m)
	}
	return c.JSON(a)
}

func signup(c *fiber.Ctx) error {
	input := gormdb.Usuarios{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	input.Id = 0
	h := sha1.New()
	h.Write([]byte(input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = i

	a := db.Create(&input)
	if a.Error != nil {
		return c.JSON(a.Error)
	}
	return c.JSON(input)
}

func forgetpassword(c *fiber.Ctx) error {

	input := gormdb.Usuarios{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	password := utils.UUID()[0:13]
	h := sha1.New()
	h.Write([]byte(password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = i

	//moduls.Send_Mail_Password(input.Email, password)

	db.Create(&input)
	return c.JSON(input)
}

func infouser(c *fiber.Ctx) error {

	//db midelware
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", "1")
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}

	return c.JSON(a)
}

func deleteuser(c *fiber.Ctx) error {
	//db midelware
	a := gormdb.Usuarios{
		Id: 11,
	}
	err2 := db.Delete(&a)
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}
	m := make(map[string]string)
	m["mensjae"] = "eliminado"

	return c.JSON(m)
}

func renuevaToken(c *fiber.Ctx) error {

	input := gormdb.Usuarios{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	password := utils.UUID()[0:13]
	h := sha1.New()
	h.Write([]byte(password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = i

	//moduls.Send_Mail_Password(input.Email, password)

	db.Create(&input)
	return c.JSON(input)
}

func users(c *fiber.Ctx) error {

	input := []gormdb.Usuarios{}
	db.Find(&input)
	return c.JSON(input)
}

func deleteById(c *fiber.Ctx) error {
	param := c.Params("id")
	temp, err := strconv.ParseUint(param, 0, 32)
	if err != nil {
		return c.JSON(err.Error)
	}
	//db midelware
	a := gormdb.Usuarios{
		Id: uint32(temp),
	}
	err2 := db.Delete(&a)
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}
	m := make(map[string]string)
	m["mensjae"] = "eliminado"

	return c.JSON(m)
}

func getById(c *fiber.Ctx) error {
	param := c.Params("id")
	temp, err := strconv.ParseUint(param, 0, 32)
	if err != nil {
		return c.JSON(err.Error)
	}

	//db midelware
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", uint32(temp))
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}

	return c.JSON(a)
}

func changepassword(c *fiber.Ctx) error {

	input := gormdb.Usuarios{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	password := utils.UUID()[0:13]
	h := sha1.New()
	h.Write([]byte(password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = i

	//moduls.Send_Mail_Password(input.Email, password)

	db.Create(&input)
	return c.JSON(input)
}

func updateuser(c *fiber.Ctx) error {

	input := gormdb.Usuarios{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	password := utils.UUID()[0:13]
	h := sha1.New()
	h.Write([]byte(password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = i

	//moduls.Send_Mail_Password(input.Email, password)

	db.Create(&input)
	return c.JSON(input)
}

/*
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}*/
