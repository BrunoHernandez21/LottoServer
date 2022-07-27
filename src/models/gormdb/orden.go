package gormdb

import "time"

type Ordenes struct {
	Id            uint32     `json:"id"`
	Status        string     `json:"status"`
	Fecha_emitido *time.Time `json:"fecha_emitido"`
	Total         int32      `json:"total"`
	Iva           *float32   `json:"iva"`
	Descuento     *float32   `json:"descuento"`
	Total_iva     *float32   `json:"total_iva"`
	Usuario_id    uint32     `json:"usuario_id"`
}

func (product *Ordenes) TableName() string {
	return "ordenes"
}
