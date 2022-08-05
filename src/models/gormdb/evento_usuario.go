package gormdb

import "time"

type Evento_usuario struct {
	Id             uint32     `json:"id"`
	Activo         *bool      `json:"activo,omitempty"`
	Fecha          *time.Time `json:"fecha,omitempty"`
	Views_count    *uint32    `json:"views_count,omitempty"`
	Like_count     *uint32    `json:"like_count,omitempty"`
	Comments_count *uint32    `json:"comments_count,omitempty"`
	Dislikes_count *uint32    `json:"dislikes_count,omitempty"`
	Saved_count    *uint32    `json:"saved_count,omitempty"`
	Shared_count   *uint32    `json:"shared_count,omitempty"`
	Usuario_id     uint32     `json:"usuario_id,omitempty"`
	Evento_id      uint32     `json:"evento_id,omitempty"`
}

func (product *Evento_usuario) TableName() string {
	return "evento_usuario"
}
