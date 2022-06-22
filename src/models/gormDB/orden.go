package gormdb

type Orden struct {
	Id           uint32   `json:"id"`
	Amount       *float32 `json:"amount,omitempty"`
	Fecha_orden  *string  `json:"fecha_orden,omitempty"`
	Id_charges   *string  `json:"id_charges,omitempty"`
	Orden_status *string  `json:"orden_status,omitempty"`
	Usuario_id   *uint32  `json:"usuario_id,omitempty"`
}
