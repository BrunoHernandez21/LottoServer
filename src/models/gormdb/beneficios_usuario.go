package gormdb

import "time"

type Beneficios_usuario struct {
	Id           uint32    `json:"id"`
	Activo       bool      `json:"activo"`
	Fecha_inicio time.Time `json:"fecha_inicio"`
	Fecha_fin    time.Time `json:"fecha_fin"`
	Usuario_id   uint32    `json:"usuario_id"`
	Beneficio_id uint32    `json:"beneficio_id"`
	Plan_id      uint32    `json:"plan_id"`
}

func (product *Beneficios_usuario) TableName() string {
	return "beneficios_usuario"
}
