package gormdb

import "time"

type Compra struct {
	Id           uint32     `json:"id"`
	Fecha_pagado *time.Time `json:"fecha_compra,omitempty"`
	Usuario_id   uint32     `json:"usuario_id,omitempty"`
	Carrito_id   uint32     `json:"carrito_id"`
}

func (product *Compra) TableName() string {
	return "compra"
}
