package gormdb

import "time"

type Apuesta_usuario struct {
	Id             uint32     `json:"id,omitempty"`
	Activo         *bool      `json:"activo,omitempty"`
	Cantidad       uint32     `json:"cantidad,omitempty"`
	Monto          uint32     `json:"monto,omitempty"`
	Fecha          *time.Time `json:"fecha,omitempty"`
	Likes          *uint32    `json:"likes,omitempty"`
	Vistas         *uint32    `json:"vistas,omitempty"`
	Apuesta        *uint32    `json:"apuesta,omitempty"`
	Usuario        *uint32    `json:"usuario,omitempty"`
	Suscribcion_id *uint32    `json:"suscribcion_id,omitempty"`
}

func (product *Apuesta_usuario) TableName() string {
	return "Apuesta_usuario"
}
