package compuestas

import "time"

type CarritoPlan struct {
	Id            uint32     `json:"id"`
	Cantidad      int32      `json:"cantidad"`
	Total_linea   *float32   `json:"total_linea"`
	Puntos_linea  *float32   `json:"puntos_linea"`
	Fecha_carrito *time.Time `json:"fecha_carrito"`
	Plan_id       uint32     `json:"plan_id"`
	Titulo        *string    `json:"titulo,omitempty"`
	Descripcion   *string    `json:"descripcion,omitempty"`
	Moneda        *string    `json:"moneda,omitempty"`
	Suscribcion   bool       `json:"suscribcion"`
}

func (product *CarritoPlan) TableName() string {
	return "carritoplan"
}
