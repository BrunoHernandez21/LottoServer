package gormdb

import "time"

type Evento struct {
	Id                  uint32     `json:"id"`
	Activo              *uint32    `json:"activo,omitempty"`
	Fechahora_evento    *time.Time `json:"fechahora_evento,omitempty"`
	Premio_cash         *float32   `json:"premio_cash,omitempty"`
	Acumulado           *float32   `json:"acumulado,omitempty"`
	Premio_otros        *string    `json:"premio_otros,omitempty"`
	Moneda              *string    `json:"moneda,omitempty"`
	Categoria_evento_id uint32     `json:"categoria_evento_id,omitempty"`
	Video_id            *string    `json:"video_id,omitempty"`
}

func (product *Evento) TableName() string {
	return "evento"
}
