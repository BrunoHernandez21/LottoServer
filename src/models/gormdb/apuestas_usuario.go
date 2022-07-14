package gormdb

import "time"

type Apuesta_usuario struct {
	Id          uint32     `json:"id,omitempty"`
	Activo      *bool      `json:"activo,omitempty"`
	Cantidad    uint32     `json:"cantidad,omitempty"`
	Fecha       *time.Time `json:"fecha,omitempty"`
	Comentarios uint32     `json:"comentarios,omitempty"`
	Likes       uint32     `json:"likes,omitempty"`
	Vistas      uint32     `json:"vistas,omitempty"`
	Dislikes    uint32     `json:"dislikes,omitempty"`
	Usuario_id  uint32     `json:"usuario_id,omitempty"`
	Apuesta_id  uint32     `json:"apuesta_id,omitempty"`
}

func (product *Apuesta_usuario) TableName() string {
	return "Apuesta_usuario"
}
