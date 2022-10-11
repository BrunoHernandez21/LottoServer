package gormdb

import "time"

type Ordenes struct {
	Id             uint32     `json:"id"`
	Status         string     `json:"status"`
	Fecha_emitido  *time.Time `json:"fecha_emitido"`
	Precio_total   int32      `json:"precio_total"`
	Puntos_total   int32      `json:"puntos_total"`
	Usuario_id     uint32     `json:"usuario_id"`
	Is_suscription bool       `json:"is_suscription"`
	Moneda         string     `json:"moneda"`
}

func (product *Ordenes) TableName() string {
	return "ordenes"
}
