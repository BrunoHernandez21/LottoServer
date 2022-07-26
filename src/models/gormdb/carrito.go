package gormdb

import "time"

type Carrito struct {
	Id              uint32     `json:"id"`
	Activo          *bool      `json:"activo"`
	Cantidad        int32      `json:"cantidad"`
	Total_linea     *float32   `json:"total_linea"`
	Precio_unitario *float32   `json:"precio_unitario"`
	Descuento       *float32   `json:"descuento"`
	Fecha_carrito   *time.Time `json:"fecha_carrito"`
	Plan_id         uint32     `json:"plan_id"`
	Usuario_id      uint32     `json:"usuario_id"`
}

func (product *Carrito) TableName() string {
	return "carrito"
}
