package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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
	fmt.Println(unikey)
	out.Codigo_referido = &unikey
	errdb = db.Create(&out)
	if errdb.Error != nil {
		m["mensaje"] = "Error en la base de datos"
		return c.Status(500).JSON(m)
	}
	errdb = db.Where("email = ?", input.Email).Find(&out)
	if errdb.Error != nil {
		m["mensaje"] = "Error en la base de datos"
		return c.Status(500).JSON(m)
	}
	if input.Referido_por != nil {
		referido := gormdb.Referido{
			User_id: out.Id,
			Codigo:  *input.Referido_por,
			Cobrado: false,
		}
		errdb = db.Create(&referido)
		if errdb.Error != nil {
			m["mensaje"] = "Error en la base de datos"
			return c.Status(500).JSON(m)
		}
	}

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

func infouser(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "id = ?", c.Locals("userID"))
	pass := ""
	a.Password = &pass
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(a)
}

func deleteuser(c *fiber.Ctx) error {
	m := make(map[string]string)
	headers := c.GetReqHeaders()
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "La cuenta no existe"
		return c.JSON(m)
	}
	h := sha1.New()
	h.Write([]byte(headers["Password"]))
	i := hex.EncodeToString(h.Sum(nil))
	var password string = i
	if *a.Password != password {
		m["mensaje"] = "Contraseña invalida"
		return c.JSON(m)
	}
	//db midelware
	errdb = db.Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	user_rol := gormdb.Usuarios_roles{}
	errdb = db.Find(&user_rol, "User_id = ?", a.Id).Delete(&user_rol)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "Eliminado satisfactoriamente"
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
	return c.JSON(inputs.Set_login{
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
	errdb := db.Find(&a, "id = ?", param).Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	user_rol := gormdb.Usuarios_roles{}
	errdb = db.Find(&user_rol, "User_id = ?", a.Id).Delete(&user_rol)
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

func changepassword(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := inputs.Get_ChangePassword{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "Datos insuficientes"
		return c.Status(500).JSON(m)
	}
	a := gormdb.Usuarios{}
	err2 := db.Find(&a, "id = ?", c.Locals("userID"))
	if err2.Error != nil {

		m["mensaje"] = "Usuario no registrado"
		return c.Status(500).JSON(m)
	}

	h := sha1.New()
	h.Write([]byte(input.Password))
	i := hex.EncodeToString(h.Sum(nil))
	a.Password = &i

	db.Save(&a)
	m["mensaje"] = "Contraseña cambiada con exito"
	return c.JSON(m)
}

func updateuser(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Usuarios{}
	errdb := db.Find(&a, "id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = "Usuario no registrado"
		return c.Status(500).JSON(m)
	}

	input := gormdb.Usuarios{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	input.Id = a.Id
	input.Activo = a.Activo
	input.Email = a.Email
	input.Password = a.Password
	if input.Apellidom == nil {
		input.Apellidom = a.Apellidom
	}
	if input.Apellidop == nil {
		input.Apellidop = a.Apellidop
	}
	if input.Fecha_nacimiento == nil {
		input.Fecha_nacimiento = a.Fecha_nacimiento
	}
	if input.Nombre == nil {
		input.Nombre = a.Nombre
	}

	errdb = db.Save(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}

func createDireccion(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Direccion{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "Datos insuficientes"
		return c.Status(500).JSON(m)
	}
	input.Id = 0
	id, ok := c.Locals("userID").(uint32)
	if ok {
		input.User_id = &id
	} else {
		m["mensaje"] = "error interno"
		return c.Status(500).JSON(m)
	}
	errdb := db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}

func updateDireccion(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Direccion{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = "Datos insuficientes"
		return c.Status(500).JSON(m)
	}
	userid, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "error interno"
		return c.Status(500).JSON(m)
	}
	compare := gormdb.Direccion{}
	errdb := db.Find(&compare, "id = ?", input.Id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if *compare.User_id != userid {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	} else {
		db.Save(input)
	}
	return c.JSON(input)
}

func getDireccion(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	outs := []gormdb.Direccion{}
	errdb := db.Find(&outs, "User_id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["direcciones"] = outs
	return c.JSON(m)
}

func deleteDireccion(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	direccion := gormdb.Direccion{}
	errdb := db.Find(&direccion, "id = ? AND user_id = ?", c.Params("id"), c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if direccion.Id == 0 {
		m["mensaje"] = "esta direccion no es tuya o no existe"
		return c.Status(500).JSON(m)
	}
	errdb = db.Delete(&direccion)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["resp"] = "eliminado correctamente"
	return c.JSON(m)
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
