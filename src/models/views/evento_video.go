package views

import "time"

type EventoVideo struct {
	Id               uint32     `json:"id"`
	Fechahora_evento *time.Time `json:"fechahora_evento"`
	Premio_cash      *float32   `json:"premio_cash"`
	Acumulado        *float32   `json:"acumulado"`
	Premio_otros     *string    `json:"premio_otros"`
	Moneda           *string    `json:"moneda"`
	Costo            uint32     `json:"costo"`
	Is_views         bool       `json:"is_views"`
	Is_like          bool       `json:"is_like"`
	Is_comments      bool       `json:"is_comments"`
	Is_saved         bool       `json:"is_saved"`
	Is_shared        bool       `json:"is_shared"`
	Is_dislikes      bool       `json:"is_dislikes"`
	Vid_id           uint32     `json:"vid_id"`
	Artista          *string    `json:"artista,omitempty"`
	Canal            *string    `json:"canal,omitempty"`
	Fecha_video      *time.Time `json:"fecha_video,omitempty"`
	Video_id         *string    `json:"video_id,omitempty"`
	Thumblary        *string    `json:"thumblary,omitempty"`
	Titulo           *string    `json:"titulo,omitempty"`
	Url_video        *string    `json:"url_video,omitempty"`
	Genero           *string    `json:"genero,omitempty"`
	Proveedor        *string    `json:"proveedor,omitempty"`
}

func (product *EventoVideo) TableName() string {
	return "eventos_videos"
}
