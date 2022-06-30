package gormdb

import "time"

type Compra struct {
	Id            uint32     `json:"id"`
	Cantidad      uint32     `json:"cantidad"`
	Amount        float32    `json:"amount"`
	Fecha_compra  *time.Time `json:"Fecha_compra,omitempty"`
	Usuario_id    *uint32    `json:"Usuario_id,omitempty"`
	Suscripcio_id uint32     `json:"suscripcio_id"`
}

func (product *Compra) TableName() string {
	return "Compra"
}
