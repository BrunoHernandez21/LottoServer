package gormdb

import "time"

type Evento_usuario struct {
	Id             uint32     `json:"id"`
	Activo         *bool      `json:"activo"`
	Fecha          *time.Time `json:"fecha"`
	Views_count    *uint64    `json:"views_count"`
	Like_count     *uint64    `json:"like_count"`
	Comments_count *uint64    `json:"comments_count"`
	Dislikes_count *uint64    `json:"dislikes_count"`
	Saved_count    *uint64    `json:"saved_count"`
	Shared_count   *uint64    `json:"shared_count"`
	Usuario_id     uint32     `json:"usuario_id"`
	Evento_id      uint32     `json:"evento_id"`
}

func (product *Evento_usuario) TableName() string {
	return "evento_usuario"
}
