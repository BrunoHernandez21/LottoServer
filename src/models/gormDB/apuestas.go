package gormdb

import "time"

type Apuestas struct {
	Id                   uint32     `json:"id"`
	Activo               bool       `json:"activo"`
	Fechahoraapuesta     *time.Time `json:"fechahoraapuesta,omitempty"`
	Precio               *float32   `json:"precio,omitempty"`
	Premio               *string    `json:"premio,omitempty"`
	Categoria_apuesta_id *uint32    `json:"categoria_apuesta_id,omitempty"`
	Tipo_apuesta_id      *uint32    `json:"tipo_apuesta_id,omitempty"`
	Video_id             *uint32    `json:"video_id,omitempty"`
}
