package gormdb

import "time"

type Carrito struct {
	Id            uint32     `json:"id"`
	Activo        *bool      `json:"activo"`
	Status        *string    `json:"status"`
	Cantidad      int32      `json:"cantidad"`
	Total         *float32   `json:"total"`
	Fecha_carrito *time.Time `json:"fecha_carrito"`
	Plan_id       uint32     `json:"plan_id"`
	Usuario_id    uint32     `json:"usuario_id"`
}

func (product *Carrito) TableName() string {
	return "carrito"
}
