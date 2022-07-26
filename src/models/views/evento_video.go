package views

import "time"

type EventoVideo struct {
	Id                  uint32     `json:"id"`
	Activo              *bool      `json:"activo"`
	Fechahora_evento    *time.Time `json:"status"`
	Premio_cash         *float32   `json:"premio_cash"`
	Acumulado           *float32   `json:"acumulado"`
	Premio_otros        *string    `json:"premio_otros"`
	Moneda              *string    `json:"moneda"`
	Categoria_evento_id uint32     `json:"categoria_evento_id"`
	Artista             *string    `json:"artista,omitempty"`
	Canal               *string    `json:"canal,omitempty"`
	Fecha_video         *time.Time `json:"fecha_video,omitempty"`
	Video_id            *string    `json:"video_id,omitempty"`
	Thumblary           *string    `json:"thumblary,omitempty"`
	Titulo              *uint32    `json:"titulo,omitempty"`
	Url_video           *string    `json:"url_video,omitempty"`
	Genero              *string    `json:"genero,omitempty"`
	Proveedor           *string    `json:"proveedor,omitempty"`
}

func (product *EventoVideo) EventoVideo() string {
	return "eventos_videos"
}
