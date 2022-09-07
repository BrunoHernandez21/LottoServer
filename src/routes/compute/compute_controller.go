package compute

import (
	"encoding/json"
	"fmt"
	"lottomusic/src/config"
	"lottomusic/src/models/gormdb"
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

	for _, evento := range eventos {
		fmt.Println(evento.Video_id)
		if evento.Video_id != nil {
			//Peticion HTTP GET
			a := fiber.AcquireAgent()
			req := a.Request()
			req.Header.SetMethod("GET")
			req.SetRequestURI(config.YTestadistics + *evento.Video_id)
			myTime := time.Now()
			if err := a.Parse(); err != nil {
				m["mensaje"] = err.Error()
				return c.Status(500).JSON(m)
			}
			code, body, _ := a.Bytes()
			if code != 200 {
				break
			}
			// parseo
			var dat map[string]interface{}
			if err := json.Unmarshal(body, &dat); err != nil {
				m["mensaje"] = err.Error()
				return c.Status(500).JSON(m)
			}
			var Like_count uint32
			var Views_count uint32
			var Comments_count uint32
			Dislikes_count := uint32(0)
			Saved_count := uint32(0)
			Shared_count := uint32(0)
			for _, item := range dat["items"].([]interface{}) {
				j, a := item.(map[string]interface{})
				if a {
					s, ah := j["statistics"].(map[string]interface{})
					if ah {
						s1, a1 := s["viewCount"].(string)
						if a1 {
							v, _ := strconv.ParseUint(s1, 10, 64)
							Views_count = uint32(v)
						}
						s2, a2 := s["commentCount"].(string)
						if a2 {
							c, _ := strconv.ParseUint(s2, 10, 64)
							Comments_count = uint32(c)
						}
						s3, a3 := s["likeCount"].(string)
						if a3 {
							l, _ := strconv.ParseUint(s3, 10, 64)
							Like_count = uint32(l)
						}
					}
				}
			}
			//agregar a la lista
			db.Create(&gormdb.VideosEstadisticas{
				Video_id:       evento.Vid_id,
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

	m["mensaje"] = "Creado con exito"
	m["time"] = time.Now().String()
	return c.JSON(m)
}

func winner(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	m["resp"] = "Hola mundo"
	return c.Status(200).JSON(m)
}
