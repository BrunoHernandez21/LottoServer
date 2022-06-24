package gormdb

type Item_compra struct {
	Id        uint32  `json:"id"`
	Cantidad  *int32  `json:"cantidad,omitempty"`
	Compra_id *uint32 `json:"compra_id,omitempty"`
	Orden_id  *uint32 `json:"orden_id,omitempty"`
	Id_plan   *uint32 `json:"id_plan,omitempty"`
}

func (product *Item_compra) TableName() string {
	return "Item_compra"
}
