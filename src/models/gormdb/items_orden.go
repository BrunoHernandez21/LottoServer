package gormdb

type ItemsOrden struct {
	Cantidad     uint32  `json:"cantidad"`
	Total_linea  float64 `json:"total_linea"`
	Puntos_linea float32 `json:"puntos_linea"`
	Orden_id     uint32  `json:"orden_id"`
	Plan_id      uint32  `json:"plan_id"`
	Moneda       string  `json:"moneda"`
}

func (product *ItemsOrden) TableName() string {
	return "items_orden"
}
