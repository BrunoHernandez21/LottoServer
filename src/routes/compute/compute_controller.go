package compute

import (
	"encoding/json"
	"fmt"
	"lottomusic/src/config"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/youtube"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// process

func process_statistics(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	videos := []gormdb.Videos{}

	errdb := db.Find(&videos, "activo = ? AND proveedor = ?", true, "Youtube")
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	toDBT := []gormdb.VideosEstadisticas{}
	for _, its := range videos {
		if its.Video_id != nil {
			//Peticion HTTP GET
			a := fiber.AcquireAgent()
			req := a.Request()
			req.Header.SetMethod("GET")
			req.SetRequestURI(config.YTestadistics + *its.Video_id)
			myTime := time.Now().Local()
			if err := a.Parse(); err != nil {
				m["mensaje"] = err.Error()
				return c.Status(500).JSON(m)
			}
			code, body, _ := a.Bytes()
			if code != 200 {
				temp := make(map[string]interface{})
				json.Unmarshal(body, &temp)
				return c.Status(500).JSON(temp)
			}
			// parseo
			var resp = youtube.YtResponse{}
			if err := json.Unmarshal(body, &resp); err != nil {
				m["mensaje"] = err.Error()
				return c.Status(500).JSON(m)
			}
			if len(resp.Items) <= 0 {
				m["mensaje"] = "yt no proporciono estadisticas"
				return c.Status(500).JSON(m)
			}

			Dislikes_count := uint32(0)
			Saved_count := uint32(0)
			Shared_count := uint32(0)

			v, _ := strconv.ParseUint(resp.Items[0].Statistics.ViewCount, 10, 64)
			Views_count := uint32(v)

			c, _ := strconv.ParseUint(resp.Items[0].Statistics.CommentCount, 10, 64)
			Comments_count := uint32(c)

			l, _ := strconv.ParseUint(resp.Items[0].Statistics.LikeCount, 10, 64)
			Like_count := uint32(l)

			//agregar a la lista
			toDBT = append(toDBT, gormdb.VideosEstadisticas{
				Video_id:       its.Id,
				Fecha:          myTime,
				Like_count:     &Like_count,
				Views_count:    &Views_count,
				Comments_count: &Comments_count,
				Dislikes_count: &Dislikes_count,
				Saved_count:    &Saved_count,
				Shared_count:   &Shared_count,
			})
		}
	}
	db.Create(&toDBT)
	m["mensaje"] = "Creado con exito"
	m["time"] = time.Now().String()
	m["nan"] = videos
	return c.JSON(m)
}

func process_users(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	var resp string
	db.Raw("CALL verificar_propiedades_usuario()").Scan(&resp)
	m["resp"] = resp
	return c.Status(200).JSON(m)
}

func process_subscriptions(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	var resp string
	db.Raw("CALL verificar_suscribciones()").Scan(&resp)
	m["resp"] = resp
	return c.Status(200).JSON(m)
}

func process_winner(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	var resp string
	db.Raw("CALL generar_ganador()").Scan(&resp)
	m["resp"] = resp
	return c.Status(200).JSON(m)
}

///////////// emit
func emit_statistics(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	//Peticion HTTP GET
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("GET")
	req.SetRequestURI("https://lotto.inclusive.com.mx/api/v1/emit/estadisticas")
	if err := a.Parse(); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	code, body, _ := a.Bytes()
	if code != 200 {
		temp := make(map[string]interface{})
		json.Unmarshal(body, &temp)
		temp["error"] = "error"
		return c.Status(500).JSON(temp)
	}
	m["resp"] = "Enviado correctamente"
	return c.Status(200).JSON(m)

}

func emit_winner(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	m["resp"] = "Hola mundo"
	return c.Status(200).JSON(m)
}

//// Webhook
func stripe_webhook(c *fiber.Ctx) error {
	// input := make(map[string]interface{})

	fmt.Println(string(c.Body()))
	out := make(map[string]interface{})
	return c.Status(200).JSON(out)
}
