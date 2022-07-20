package gormdb

import "time"

type Suscripciones struct {
	Id            uint32     `json:"id"`
	Activo        *bool      `json:"activo"`
	Monto_mensual *float32   `json:"monto_mensual"`
	Fecha_create  *time.Time `json:"fecha_create"`
	Fecha_inicio  *time.Time `json:"fecha_inicio"`
	Fecha_fin     *time.Time `json:"fecha_fin"`
	Fecha_cobro   *time.Time `json:"fecha_cobro"`
	Fecha_corte   uint32     `json:"fecha_corte"`
	Tipo          string     `json:"tipo"`
	Plan_id       uint32     `json:"plan_id"`
	Usuario_id    uint32     `json:"usuario_id"`
}

func (product *Suscripciones) TableName() string {
	return "suscripciones"
}
