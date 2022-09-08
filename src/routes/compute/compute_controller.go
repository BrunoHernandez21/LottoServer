package compute

import (
	"encoding/json"
	"lottomusic/src/config"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"lottomusic/src/models/views"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func statistics(c *fiber.Ctx) error {
	m := make(map[string]string)
	eventos := []views.EventoVideo{}

	errdb := db.Find(&eventos)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	toDBT := []gormdb.VideosEstadisticas{}
	for _, its := range eventos {
		if its.Video_id != nil {
			//Peticion HTTP GET
			a := fiber.AcquireAgent()
			req := a.Request()
			req.Header.SetMethod("GET")
			req.SetRequestURI(config.YTestadistics + *its.Video_id)
			myTime := time.Now().UTC()
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
			var resp = inputs.YtResponse{}
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
				Video_id:       its.Vid_id,
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
	return c.JSON(m)
}

func emit(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	//Peticion HTTP GET
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("GET")
	req.SetRequestURI("http://187.213.68.250:25567/api/v1/emit/send/message")
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
	m["resp"] = "Enviado correctamente"
	return c.Status(200).JSON(m)

}

func winner(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	m["resp"] = "Hola mundo"
	return c.Status(200).JSON(m)
}
