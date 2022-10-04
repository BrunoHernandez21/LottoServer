package evento

import (
	"encoding/json"
	"lottomusic/src/config"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/inputs"
	"lottomusic/src/models/youtube"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func crear(c *fiber.Ctx) error {
	m := make(map[string]string)
	input := inputs.Get_Event_Video{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	// fase de comprovacion de entrada
	if input.Video_id == "" {
		m["mensaje"] = "video id no puede ser null"
		return c.Status(500).JSON(m)
	}

	if input.Evento.Costo == 0 {
		m["mensaje"] = "costo no puede ser 0"
		return c.Status(500).JSON(m)
	}
	if input.Evento.Premio_otros == nil && *input.Evento.Premio_cash == 0 {
		m["mensaje"] = "Premio no puede ser nil"
		return c.Status(500).JSON(m)
	}

	if *input.Evento.Premio_cash != 0 && input.Evento.Moneda == nil {
		m["mensaje"] = "si el premio es en cash moneda es necesario"
		return c.Status(500).JSON(m)
	}

	if !(input.Evento.Is_comments || input.Evento.Is_dislikes || input.Evento.Is_like || input.Evento.Is_saved || input.Evento.Is_shared || input.Evento.Is_views) {

		m["mensaje"] = "Selecciona los tipos permitidos para el evento"
		return c.Status(500).JSON(m)
	}
	acumulado := 0.0
	input.Evento.Acumulado = &acumulado
	input.Evento.Id = 0

	video := gormdb.Videos{}
	errdb := db.Find(&video, "Video_id = ?", input.Video_id)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	myTime := time.Now().Local()
	// Crear video si no existe
	if video.Id == 0 {
		a := fiber.AcquireAgent()
		req := a.Request()
		req.Header.SetMethod("GET")
		req.SetRequestURI(config.YTbyID + input.Video_id)

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
		ytvideo := resp.Items[0].Snippet
		if a == nil {
			m["mensaje"] = "yt no proporciono estadisticas"
			return c.Status(500).JSON(m)
		}
		activo := true
		urlvideo := "https://www.youtube.com/watch?v=" + input.Video_id
		proveedor := "Youtube"
		videonew := gormdb.Videos{
			Id:          0,
			Activo:      &activo,
			Artista:     &ytvideo.ChannelTitle,
			Canal:       &ytvideo.ChannelTitle,
			Fecha_video: &myTime,
			Video_id:    &input.Video_id,
			Titulo:      &ytvideo.Title,
			Url_video:   &urlvideo,
			Thumblary:   &ytvideo.Thumbnails.Default.URL,
			Genero:      input.Genero,
			Proveedor:   &proveedor,
		}
		//agregar a la lista
		errdb = db.Create(&videonew)
		if errdb.Error != nil {
			m["mensaje"] = errdb.Error.Error()
			return c.Status(500).JSON(m)
		}
		input.Evento.Video_id = videonew.Id
	} else {
		if !*video.Activo {
			activo := true
			video.Activo = &activo
			video.Fecha_video = &myTime
			errdb = db.Save(&video)
			if errdb.Error != nil {
				m["mensaje"] = errdb.Error.Error()
				return c.Status(500).JSON(m)
			}
		}
		input.Evento.Video_id = video.Id

	}

	var evento *gormdb.Eventos = input.Evento
	errdb = db.Create(&evento)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(evento)
}
func byid(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	a := gormdb.Eventos{}
	errdb := db.Find(&a, "id = ?", param)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "El evento no existe"
		return c.Status(500).JSON(m)
	}

	return c.JSON(a)
}
func editar(c *fiber.Ctx) error {
	m := make(map[string]string)

	input := gormdb.Eventos{}
	if err := c.BodyParser(&input); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	if input.Id == 0 {
		m["mensaje"] = "Id es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Fechahora_evento == nil {
		m["mensaje"] = "Fechahora_evento es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Premio_cash == nil && input.Premio_otros == nil {
		m["mensaje"] = "Premio_cash o Premio_otros es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Premio_cash == nil && input.Moneda == nil {
		m["mensaje"] = "moneda no puede ser nul si el premio es cash"
		return c.Status(500).JSON(m)
	}
	if input.Video_id == 0 {
		m["mensaje"] = "video_id es requerido"
		return c.Status(500).JSON(m)
	}
	if input.Costo == 0 {
		m["mensaje"] = "costo no puede ser nil"
		return c.Status(500).JSON(m)
	}

	errdb := db.Save(&input)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	return c.JSON(input)
}
func eliminar(c *fiber.Ctx) error {
	m := make(map[string]string)
	param := c.Params("id")
	//db midelware
	a := gormdb.Eventos{}
	errdb := db.Find(&a, "id = ?", param).Delete(&a)
	if errdb.Error != nil {
		m["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(m)
	}
	if a.Id == 0 {
		m["mensaje"] = "El evento no existe"
		return c.Status(500).JSON(m)
	}

	m["mensaje"] = "Eliminado Satisfactoriamente"
	return c.JSON(m)
}

func listarActivos(c *fiber.Ctx) error {
	resp := make(map[string]interface{})
	pagt := c.Params("pag")
	sizepaget := c.Params("sizepage")
	input := []gormdb.Eventos{}
	errdb := db.Find(&input, "Activo = ?", true)
	if errdb.Error != nil {
		resp["mensaje"] = errdb.Error.Error()
		return c.Status(500).JSON(resp)
	}
	page, err4 := strconv.ParseUint(pagt, 0, 32)
	if err4 != nil {
		resp["mensaje"] = err4.Error()
		return c.Status(500).JSON(resp)
	}
	sizepage, err5 := strconv.ParseUint(sizepaget, 10, 32)
	if err5 != nil {
		resp["mensaje"] = err5.Error()
		return c.Status(500).JSON(resp)
	}

	resp["pags"] = math.Round(float64(len(input)) / float64(sizepage))
	resp["pag"] = page
	init := (page - 1) * sizepage
	end := (page * sizepage) - 1
	if int(end) > len(input) {
		end = uint64(len(input))
	}
	if init > end {
		resp["videos"] = nil
	} else {
		resp["videos"] = input[init:end]
	}

	return c.JSON(resp)
}
