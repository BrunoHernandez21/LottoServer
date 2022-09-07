package gormdb

import "time"

type Suscripciones struct {
	Id            uint32     `json:"id"`
	Monto_mensual *float32   `json:"monto_mensual"`
	Fecha_creado  *time.Time `json:"fecha_creado"`
	Fecha_inicio  *time.Time `json:"fecha_inicio"`
	Fecha_fin     *time.Time `json:"fecha_fin"`
	Fecha_cobro   *time.Time `json:"fecha_cobro"`
	Dia_corte     uint32     `json:"dia_corte"`
	Tipo          string     `json:"tipo"`
	Plan_id       uint32     `json:"plan_id"`
	Usuario_id    uint32     `json:"usuario_id"`
}

func (product *Suscripciones) TableName() string {
	return "suscripciones"
}
