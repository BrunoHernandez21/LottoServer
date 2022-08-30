package gormdb

type Items_orden struct {
	Cantidad    int32    `json:"cantidad"`
	Total_linea *float32 `json:"total_linea"`
	Moneda      string   `json:"moneda"`
	Plan_id     uint32   `json:"plan_id"`
	Orden_id    uint32   `json:"orden_id"`
}

func (product *Items_orden) TableName() string {
	return "items_orden"
}
