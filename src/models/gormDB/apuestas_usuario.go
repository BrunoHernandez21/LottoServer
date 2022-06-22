package gormdb

import "time"

type Apuesta_usuario struct {
	Id             uint32     `json:"id,omitempty"`
	Comentarios    *uint32    `json:"comentarios,omitempty"`
	Fecha_creacion *time.Time `json:"fecha_creacion,omitempty"`
	Likes          *uint32    `json:"likes,omitempty"`
	Vistas         *uint32    `json:"vistas,omitempty"`
	Apuesta        *uint32    `json:"apuesta,omitempty"`
	Usuario        *uint32    `json:"usuario,omitempty"`
}
