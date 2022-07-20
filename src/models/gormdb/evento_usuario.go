package gormdb

import "time"

type Evento_usuario struct {
	Id         uint32     `json:"id"`
	Activo     *bool      `json:"activo,omitempty"`
	Fecha      *time.Time `json:"fecha,omitempty"`
	Views      *uint32    `json:"views,omitempty"`
	Comments   *uint32    `json:"comments,omitempty"`
	Like       *uint32    `json:"like,omitempty"`
	Dislikes   *uint32    `json:"dislikes,omitempty"`
	Saved      *uint32    `json:"saved,omitempty"`
	Shared     *uint32    `json:"shared,omitempty"`
	Usuario_id uint32     `json:"usuario_id,omitempty"`
	Evento_id  uint32     `json:"evento_id,omitempty"`
}

func (product *Evento_usuario) TableName() string {
	return "evento_usuario"
}
