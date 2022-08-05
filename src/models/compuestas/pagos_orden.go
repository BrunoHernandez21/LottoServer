package compuestas

import "time"

type Pagos_orden struct {
	Id                uint32     `json:"id"`
	Fecha_pagado      *time.Time `json:"fecha_pagado"`
	Stripe_id         string     `json:"stripe_id"`
	Status            string     `json:"status"`
	Fecha_emitido     *time.Time `json:"fecha_emitido"`
	Impuesto          *float32   `json:"impuesto"`
	Sub_total         *float32   `json:"sub_total"`
	Descuento_orden   *float32   `json:"descuento_orden"`
	Total             *float32   `json:"total"`
	Payment_method_id *float32   `json:"payment_method_id"`
}

func (product *Pagos_orden) TableName() string {
	return "pagos_orden"
}
