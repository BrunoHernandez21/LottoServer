package gormdb

import "time"

type Propiedades_usuario struct {
	Id                uint32    `json:"id"`
	Nivel_acceso      string    `json:"nivel_acceso"`
	Custom_attributes string    `json:"custom_attributes"`
	Fecha_inicio      time.Time `json:"fecha_inicio"`
	Fecha_fin         time.Time `json:"fecha_fin"`
	Usuario_id        uint32    `json:"usuario_id"`
}

func (product *Propiedades_usuario) TableName() string {
	return "propiedades_usuarios"
}
