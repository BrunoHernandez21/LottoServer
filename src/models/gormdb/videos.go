package gormdb

import "time"

type Videos struct {
	Id          uint32     `json:"id"`
	Activo      *bool      `json:"activo"`
	Artista     *string    `json:"artista,omitempty"`
	Canal       *string    `json:"canal,omitempty"`
	Fecha_video *time.Time `json:"fecha_video,omitempty"`
	Video_id    *string    `json:"video_id,omitempty"`
	Titulo      *string    `json:"titulo,omitempty"`
	Url_video   *string    `json:"url_video,omitempty"`
	Thumblary   *string    `json:"thumblary,omitempty"`
	Genero      *string    `json:"genero,omitempty"`
	Proveedor   *string    `json:"proveedor,omitempty"`
}

func (product *Videos) TableName() string {
	return "videos"
}
