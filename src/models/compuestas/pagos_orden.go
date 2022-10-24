package compuestas

import "time"

type Pagos_orden struct {
	Id            uint32     `json:"id"`
	Fecha_pagado  *time.Time `json:"fecha_pagado"`
	Is_error      bool       `json:"is_error"`
	Status        string     `json:"status"`
	Fecha_emitido *time.Time `json:"fecha_emitido"`
	Precio_total  *float32   `json:"precio_total"`
	Puntos_total  *float32   `json:"puntos_total"`
}

func (product *Pagos_orden) TableName() string {
	return "pagos_orden"
}
