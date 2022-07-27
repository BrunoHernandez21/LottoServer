package compuestas

import "time"

type Pagos_orden struct {
	Id            uint32     `json:"id"`
	Status        string     `json:"status"`
	Fecha_emitido *time.Time `json:"fecha_emitido"`
	Total         *float32   `json:"total"`
	Iva           *float32   `json:"iva"`
	Descuento     *float32   `json:"descuento"`
	Total_iva     *float32   `json:"total_iva"`
	Fecha_pagado  *time.Time `json:"fecha_pagado"`
	Orden_id      uint32     `json:"orden_id"`
	Stripe_id     string     `json:"stripe_id"`
}

func (product *Pagos_orden) TableName() string {
	return "pagos_orden"
}
