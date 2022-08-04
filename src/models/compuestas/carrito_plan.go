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
	Puntos          uint32     `json:"puntos"`
	Nombre          *string    `json:"nombre,omitempty"`
	Precio          *float32   `json:"precio,omitempty"`
	Moneda          string     `json:"moneda,omitempty"`
	Duracion_dias   int32      `json:"duracion_dias,omitempty"`
	Suscribcion     bool       `json:"suscribcion,omitempty"`
}

func (product *CarritoPlan) TableName() string {
	return "carritoplan"
}
