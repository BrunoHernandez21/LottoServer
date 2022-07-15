package gormdb

import "time"

type Apuesta_usuario struct {
	Id          uint32     `json:"id"`
	Activo      *bool      `json:"activo"`
	Cantidad    uint32     `json:"cantidad"`
	Fecha       *time.Time `json:"fecha,omitempty"`
	Comentarios uint32     `json:"comentarios,omitempty"`
	Likes       uint32     `json:"likes,omitempty"`
	Vistas      uint32     `json:"vistas,omitempty"`
	Dislikes    uint32     `json:"dislikes,omitempty"`
	Usuario_id  uint32     `json:"usuario_id"`
	Apuesta_id  uint32     `json:"apuesta_id"`
}

func (product *Apuesta_usuario) TableName() string {
	return "apuesta_usuario"
}
