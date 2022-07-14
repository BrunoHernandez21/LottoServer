package gormdb

import (
	"time"
)

type Orden struct {
	Id           uint32     `json:"id"`
	Activa       *bool      `json:"activa"`
	Cantidad     uint32     `json:"cantidad"`
	Amount       *float32   `json:"amount,omitempty"`
	Fecha_orden  *time.Time `json:"fecha_orden,omitempty"`
	Id_charges   *string    `json:"id_charges,omitempty"`
	Orden_status *string    `json:"orden_status,omitempty"`
	Id_plan      uint32     `json:"id_plan,omitempty"`
	Usuario_id   uint32     `json:"usuario_id,omitempty"`
}

func (product *Orden) TableName() string {
	return "orden"
}
