package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"lottomusic/src/globals"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"lottomusic/src/modules/email"
	"lottomusic/src/modules/jwts"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	m := make(map[string]string)
	//catch body
	input := inputs.Get_Login{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if (input.Email == nil) || (input.Password == nil) {
		m["mensaje"] = "informacion insuficiente"
		return c.Status(400).JSON(m)
	}
	if (*input.Email == "") || (*input.Password == "") {
		m["mensaje"] = "informacion insuficiente"
		return c.Status(400).JSON(m)
	}
	//db midelware
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "email = ?", input.Email)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "Usuario no registrado"
		return c.Status(500).JSON(m)
	}
	//password midelware
	h := sha1.New()
	h.Write([]byte(*input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = &i
	if *a.Password != *input.Password {
		m["mensaje"] = "Contraseña invalida"
		return c.JSON(m)
	}
	//// JWT midelware
	token, expireAt := jwts.GenerateToken(a.Id)
	tokentipe := "Bearer"
	rsponse := inputs.Set_login{
		User_id:      a.Id,
		Access_token: token,
		Token_type:   &tokentipe,
		Expires_in:   &expireAt,
	}

	return c.JSON(rsponse)
}

func signup(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := inputs.Get_signup{}
	err := c.BodyParser(&input)
	if err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if (input.Email == nil) || (input.Password == nil) {
		m["mensaje"] = "Email y contraseña son obligatorios"
		return c.Status(400).JSON(m)
	}

	out := gormdb.Usuarios{
		Email: input.Email,
	}
	errdb := db.Find(&out, "email = ?", out.Email)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if out.Id != 0 {
		m["mensaje"] = "El usuario ya esta registrado"
		return c.Status(500).JSON(m)
	}

	h := sha1.New()
	h.Write([]byte(*input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	var activo bool = true
	out.Activo = &activo
	out.Password = &i
	var unikey string = randomString(10)
	out.Codigo_referido = &unikey
	errdb = db.Create(&out)
	if errdb.Error != nil {
		m["mensaje"] = "Error en la base de datos"
		return c.Status(500).JSON(m)
	}
	/*TODO: creacion de referido
	Esto tambien deberian encargarse la base de datos ?*/
	if input.Referido_por != nil {
		referido := gormdb.Referido{
			Usuario_id: out.Id,
			Codigo:     *input.Referido_por,
			Cobrado:    false,
		}
		errdb = db.Create(&referido)
		if errdb.Error != nil {
			m["mensaje"] = "Error en la base de datos"
			return c.Status(500).JSON(m)
		}
	}
	out.Password = nil
	return c.JSON(out)
}

func forgetpassword(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := inputs.Get_forgetpassword{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "email = ?", input.Email)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "Usuario no registrado"
		return c.Status(500).JSON(m)
	}

	password := randomString(12)
	h := sha1.New()
	h.Write([]byte(password))
	i := hex.EncodeToString(h.Sum(nil))
	a.Password = &i
	email.Send_Recovery_Password(input.Email, password)

	errdb = db.Save(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["resp"] = "Se a enviado un correo a su cuenta"
	return c.JSON(m)
}

func renewToken(c *fiber.Ctx) error {

	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		return c.JSON(errdb.Error)
	}
	token, expireAt := jwts.GenerateToken(a.Id)
	tipe := "Bearer"
	return c.JSON(inputs.Set_login{
		Access_token: token,
		Token_type:   &tipe,
		Expires_in:   &expireAt,
	})
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var acumulado string

	for i := 0; i < length; i++ {
		acumulado += randInt(0, len(globals.ASCII_imprimibles))
	}
	return acumulado
}
func randInt(min int, max int) string {
	return globals.ASCII_imprimibles[rand.Intn(max-min)]
}

// root

func users(c *fiber.Ctx) error {
	input := []gormdb.Usuarios{}
	db.Find(&input)
	return c.JSON(input)
}

func deleteById(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "id = ?", param).Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}

func getById(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "id = ?", param)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	return c.JSON(a)
}
