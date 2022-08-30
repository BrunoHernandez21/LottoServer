package gormdb

import "time"

type Carrito struct {
	Id              uint32     `json:"id"`
	Activo          *bool      `json:"activo"`
	Cantidad        uint32     `json:"cantidad"`
	Precio_unitario *float32   `json:"precio_unitario"`
	Total_linea     *float32   `json:"total_linea"`
	Puntos_unitario *uint32    `json:"puntos_unitario"`
	Puntos_linea    *uint32    `json:"puntos_linea"`
	Moneda          *string    `json:"moneda"`
	Fecha_carrito   *time.Time `json:"fecha_carrito"`
	Plan_id         uint32     `json:"plan_id"`
	Usuario_id      uint32     `json:"usuario_id"`
}

func (product *Carrito) TableName() string {
	return "carrito"
}
