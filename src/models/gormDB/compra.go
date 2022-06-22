package gormdb

import "time"

type Compra struct {
	Id           uint32     `json:"id"`
	Activo       bool       `json:"activo"`
	Fecha_compra *time.Time `json:"Fecha_compra,omitempty"`
	Usuario_id   *uint32    `json:"Usuario_id,omitempty"`
}
