package juegos

import (
	"crypto/sha1"
	"encoding/hex"
	gormdb "lottomusic/src/models/gormDB"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func signup(c *fiber.Ctx) error {

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

func login(c *fiber.Ctx) error {

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

func logout(c *fiber.Ctx) error {

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

func infoUser(c *fiber.Ctx) error {

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
