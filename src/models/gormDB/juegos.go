package gormdb

import "time"

type Juegos struct {
	Id           uint32     `json:"id"`
	Activo       bool       `json:"activo"`
	Fecha_compra *time.Time `json:"fecha_compra,omitempty"`
	Usuario_id   *uint32    `json:"usuario_id,omitempty"`
}

func (product *Juegos) TableName() string {
	return "Juegos"
}
