package compuestas

import "time"

type CarritoPlan struct {
	Id              uint32     `json:"id"`
	Cantidad        int32      `json:"cantidad"`
	Total_linea     *float32   `json:"total_linea"`
	Precio_unitario *float32   `json:"precio_unitario"`
	Descuento       *float32   `json:"descuento"`
	Fecha_carrito   *time.Time `json:"fecha_carrito"`
	Plan_id         uint32     `json:"plan_id"`
	Cash            uint32     `json:"cash"`
	Nombre          *string    `json:"nombre,omitempty"`
	Precio          *float32   `json:"precio,omitempty"`
	Moneda          string     `json:"moneda,omitempty"`
	Suscribcion     bool       `json:"suscribcion,omitempty"`
}

func (product *CarritoPlan) TableName() string {
	return "carritoplan"
}
