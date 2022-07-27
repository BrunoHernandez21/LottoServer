package gormdb

type Items_orden struct {
	Cantidad        int32    `json:"cantidad"`
	Total_linea     *float32 `json:"total_linea"`
	Precio_unitario *float32 `json:"precio_unitario"`
	Moneda          string   `json:"moneda"`
	Descuento       *float32 `json:"Descuento"`
	Plan_id         uint32   `json:"Plan_id"`
	Orden_id        uint32   `json:"orden_id"`
}

func (product *Items_orden) TableName() string {
	return "items_orden"
}
