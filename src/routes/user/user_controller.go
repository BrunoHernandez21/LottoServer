package user

import (
	"crypto/sha1"
	"encoding/hex"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"

	"github.com/gofiber/fiber/v2"
)

// direction
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
		input.Usuario_id = &id
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
	input.Usuario_id = &userid
	compare := gormdb.Direccion{}
	errdb := db.Find(&compare, "id = ?", input.Id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if *compare.Usuario_id != userid {
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
	errdb := db.Find(&outs, "Usuario_id = ?", c.Locals("userID"))
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
	errdb := db.Find(&direccion, "id = ? AND Usuario_id = ?", c.Params("id"), c.Locals("userID"))
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

////// payment-method

func createTarjeta(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := gormdb.Payment_method{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	input.Id = 0
	input.Usuario_id = userID
	input.Activo = true

	errdb := db.Create(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func editTarjeta(c *fiber.Ctx) error {

	m := make(map[string]string)
	input := gormdb.Payment_method{}
	compare := gormdb.Payment_method{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	input.Usuario_id = userID
	errdb := db.Find(&compare, "Id = ? ", input.Id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if compare.Id == 0 {
		m["mensaje"] = "No exite la tarjeta"
		return c.Status(500).JSON(m)
	}
	if compare.Id != input.Id {
		m["mensaje"] = "No te pertenece la tarjeta"
		return c.Status(500).JSON(m)
	}

	input.Usuario_id = userID
	input.Activo = true

	errdb = db.Save(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func deleteTarjeta(c *fiber.Ctx) error {
	m := make(map[string]string)
	a := gormdb.Payment_method{}
	err := db.Find(&a, "id = ?", c.Params("id"))
	if (err.Error != nil) || (a.Id == 0) {
		m["mensaje"] = "tarjeta no encontrado"
		return c.Status(500).JSON(m)
	}
	userID, ok := c.Locals("userID").(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}
	if a.Usuario_id != userID {
		m["mensaje"] = "Esta tarjeta no es tuya"
		return c.Status(500).JSON(m)
	}
	errdb := db.Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["mensaje"] = "Eliminado con exito"
	return c.JSON(m)
}

func listarTarjeta(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := []gormdb.Payment_method{}
	errdb := db.Find(&input, "usuario_id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["tarjetas"] = input
	return c.JSON(m)

}

// user-state

func propiedades(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	user_pro := gormdb.Propiedades_usuario{}
	errdb := db.Find(&user_pro, "usuario_id = ?", c.Locals("userID"))
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["propiedades"] = user_pro
	return c.JSON(m)
}

func suscripcion(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	input := gormdb.Suscripciones{}
	err := db.Find(&input, "usuario_id = ?", c.Locals("userID"))
	if err.Error != nil {
		m["mensaje"] = err.Error.Error()
		return c.Status(500).JSON(m)
	}
	m["suscripcion"] = input
	return c.JSON(m)
}

func cartera(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	var user = c.Locals("userID")
	userID, ok := user.(uint32)
	if !ok {
		m["mensaje"] = "internal error"
		return c.Status(500).JSON(m)
	}

	cartera := gormdb.Carteras{}
	errdb := db.Find(&cartera, "Usuario_id = ?", userID)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if cartera.Id == 0 {
		cartera.Usuario_id = userID
		errdb = db.Create(&cartera)
		if errdb.Error != nil {
			m["mensaje"] = errdb.Error.Error()
			return c.Status(500).JSON(m)
		}
	}
	errdb = db.Find(&cartera, "Usuario_id = ?", userID)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}

	return c.JSON(cartera)
}

// user

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

	m["mensaje"] = "Eliminado satisfactoriamente"
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
	input.Password = nil
	return c.JSON(input)
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
