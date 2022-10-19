package gormdb

type ItemsOrden struct {
	Cantidad     int32    `json:"cantidad"`
	Total_linea  *float32 `json:"total_linea"`
	Puntos_linea int32    `json:"puntos_linea"`
	Moneda       string   `json:"moneda"`
	Plan_id      uint32   `json:"plan_id"`
	Orden_id     uint32   `json:"orden_id"`
	Titulo       *string  `json:"titulo"`
}

func (product *ItemsOrden) TableName() string {
	return "items_orden"
}
