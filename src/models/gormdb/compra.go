package gormdb

import "time"

type Compra struct {
	Id           uint32     `json:"id"`
	Cantidad     uint32     `json:"cantidad"`
	Amount       float32    `json:"amount"`
	Fecha_compra *time.Time `json:"fecha_compra,omitempty"`
	Usuario_id   uint32     `json:"usuario_id,omitempty"`
	Plan_id      uint32     `json:"plan_id"`
}

func (product *Compra) TableName() string {
	return "compra"
}
