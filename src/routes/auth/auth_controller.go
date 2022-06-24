package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"lottomusic/src/models/auth"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/modules/email"
	"lottomusic/src/modules/jwts"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func login(c *fiber.Ctx) error {
	m := make(map[string]string)
	//catch body
	input := auth.Get_Login{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if (input.Username == nil) || (input.Password == nil) {
		m["mensjae"] = "informacion insuficiente"
		return c.JSON(m)
	}
	//db midelware
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "email = ?", input.Username)
	if errdb.Error != nil {
		return c.JSON(errdb.Error)
	}
	if a.Id == 0 {
		m["mensjae"] = "Usuario no registrado"
		return c.JSON(m)
	}
	//password midelware
	h := sha1.New()
	h.Write([]byte(*input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = &i
	if a.Password != *input.Password {
		m["mensjae"] = "Contraseña invalida"
		return c.JSON(m)
	}
	//// JWT midelware
	token, expireAt := jwts.GenerateToken(a.Id)
	tokentipe := "Bearer"
	rsponse := auth.Set_login{
		Access_token: token,
		Token_type:   &tokentipe,
		Expires_in:   &expireAt,
	}

	return c.JSON(rsponse)
}

func signup(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Usuarios{}
	err := c.BodyParser(&input)
	if err != nil {
		return c.JSON(err)
	}
	if (input.Email == "") || (input.Password == "") {
		m["mensjae"] = "informacion insuficiente"
		return c.JSON(m)
	}
	input.Id = 0
	input.Activo = true
	h := sha1.New()
	h.Write([]byte(input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	input.Password = i

	errdb := db.Create(&input)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}
	errdb = db.Find(&input, "email = ?", input.Email)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}
	user_rol := gormdb.Usuarios_roles{
		User_id: input.Id,
		Role_id: 1,
	}
	errdb = db.Create(&user_rol)
	if errdb.Error != nil {
		return c.JSON(errdb)
	}

	return c.JSON(input)
}

func forgetpassword(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := auth.Get_forgetpassword{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "email = ?", input.Email)
	if errdb.Error != nil {
		m["mensjae"] = "Usuario no registrado"
		return c.JSON(m)
	}

	password := utils.UUID()[0:13]
	h := sha1.New()
	h.Write([]byte(password))
	i := hex.EncodeToString(h.Sum(nil))
	a.Password = i
	email.Send_Recovery_Password(input.Email, password)

	errdb = db.Save(&a)
	if errdb.Error != nil {
		return c.JSON(errdb.Error)
	}
	m["mensaje"] = "Se a enviado un correo a su cuenta"
	return c.JSON(m)
}

func infouser(c *fiber.Ctx) error {
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", c.Locals("userID"))
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}
	return c.JSON(a)
}

func deleteuser(c *fiber.Ctx) error {
	m := make(map[string]string)
	headers := c.GetReqHeaders()
	a := gormdb.Usuarios{}
	err := db.Find(&a, "id = ?", c.Locals("userID"))
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	if a.Id == 0 {
		m["mensjae"] = "La cuenta no existe"
		return c.JSON(m)
	}
	h := sha1.New()
	h.Write([]byte(headers["Password"]))
	i := hex.EncodeToString(h.Sum(nil))
	var password string = i
	if a.Password != password {
		m["mensjae"] = "Contraseña invalida"
		return c.JSON(m)
	}
	//db midelware
	err = db.Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	user_rol := gormdb.Usuarios_roles{}
	err = db.Find(&user_rol, "User_id = ?", a.Id).Delete(&user_rol)
	if err.Error != nil {
		return c.JSON(err.Error)
	}

	m["mensjae"] = "Eliminado satisfactoriamente"
	return c.JSON(m)
}

func renuevaToken(c *fiber.Ctx) error {

	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		return c.JSON(errdb.Error)
	}
	token, expireAt := jwts.GenerateToken(a.Id)
	tipe := "Bearer"
	return c.JSON(auth.Set_login{
		Access_token: token,
		Token_type:   &tipe,
		Expires_in:   &expireAt,
	})
}

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
	err := db.Find(&a, "id = ?", param).Delete(&a)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	user_rol := gormdb.Usuarios_roles{}
	err = db.Find(&user_rol, "User_id = ?", a.Id).Delete(&user_rol)
	if err.Error != nil {
		return c.JSON(err.Error)
	}
	m["mensjae"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}

func getById(c *fiber.Ctx) error {
	param := c.Params("id")
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", param)
	if err2.Error != nil {
		return c.JSON(err2.Error)
	}

	return c.JSON(a)
}

func changepassword(c *fiber.Ctx) error {

	input := auth.Get_ChangePassword{}
	if err := c.BodyParser(&input); err != nil {
		m := make(map[string]string)
		m["mensjae"] = "Datos insuficientes"
		return c.JSON(m)
	}
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", c.Locals("userID"))
	if err2.Error != nil {
		m := make(map[string]string)
		m["mensjae"] = "Usuario no registrado"
		return c.JSON(m)
	}

	h := sha1.New()
	h.Write([]byte(input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	a.Password = i

	db.Save(&a)

	return c.JSON(auth.Set_ChangePassword{
		Mensaje: "Contraseña cambiada con exito",
	})
}

func updateuser(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", c.Locals("userID"))
	if err2.Error != nil {
		m["mensjae"] = "Usuario no registrado"
		return c.JSON(m)
	}

	input := gormdb.Usuarios{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	input.Id = a.Id
	input.Activo = a.Activo
	input.Email = a.Email
	input.Password = a.Password
	if input.Apellidom == "" {
		input.Apellidom = a.Apellidom
	}
	if input.Apellidop == "" {
		input.Apellidop = a.Apellidop
	}
	if input.Fecha_nacimiento == "" {
		input.Fecha_nacimiento = a.Fecha_nacimiento
	}
	if input.Nombre == "" {
		input.Nombre = a.Nombre
	}

	db.Save(&input)
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
