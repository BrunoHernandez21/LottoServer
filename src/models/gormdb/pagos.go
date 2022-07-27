package gormdb

import "time"

type Pagos struct {
	Id           uint32     `json:"id"`
	Fecha_pagado *time.Time `json:"fecha_compra,omitempty"`
	Usuario_id   uint32     `json:"usuario_id,omitempty"`
	Orden_id     uint32     `json:"orden_id,omitempty"`
	Stripe_id    string     `json:"stripe_id"`
}

func (product *Pagos) TableName() string {
	return "pagos"
}
