package gormdb

import "time"

type Videos struct {
	Id          uint32     `json:"id"`
	Activo      *bool      `json:"activo"`
	Artista     *string    `json:"Artista,omitempty"`
	Canal       *string    `json:"Canal,omitempty"`
	Fecha_video *time.Time `json:"Fecha_video,omitempty"`
	Id_video    *string    `json:"Id_video,omitempty"`
	Titulo      *string    `json:"Titulo,omitempty"`
	Url_video   *string    `json:"Url_video,omitempty"`
	Thumblary   *string    `json:"thumblary,omitempty"`
}

func (product *Videos) TableName() string {
	return "Videos"
}
