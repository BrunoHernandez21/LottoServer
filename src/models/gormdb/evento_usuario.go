package gormdb

import "time"

type Evento_usuario struct {
	Id          uint32     `json:"id"`
	Activo      *uint32    `json:"activo,omitempty"`
	Fecha       *time.Time `json:"fecha,omitempty"`
	Comentarios *uint32    `json:"comentarios,omitempty"`
	Likes       *uint32    `json:"likes,omitempty"`
	Vistas      *uint32    `json:"vistas,omitempty"`
	Shared      *uint32    `json:"shared,omitempty"`
	Usuario_id  *uint32    `json:"usuario_id,omitempty"`
	Evento_id   *uint32    `json:"evento_id,omitempty"`
}

func (product *Evento_usuario) TableName() string {
	return "evento_usuario"
}
