package gormdb

import "time"

type Eventos struct {
	Id               uint32     `json:"id"`
	Activo           bool       `json:"activo,omitempty"`
	Fechahora_evento *time.Time `json:"fechahora_evento,omitempty"`
	Premio_cash      *float32   `json:"premio_cash,omitempty"`
	Acumulado        *float32   `json:"acumulado,omitempty"`
	Premio_otros     *string    `json:"premio_otros,omitempty"`
	Moneda           *string    `json:"moneda,omitempty"`
	Video_id         uint32     `json:"video_id,omitempty"`
	Costo            uint32     `json:"costo,omitempty"`
	Is_views         bool       `json:"is_views,omitempty"`
	Is_like          bool       `json:"is_like,omitempty"`
	Is_comments      bool       `json:"is_comments,omitempty"`
	Is_saved         bool       `json:"is_saved,omitempty"`
	Is_shared        bool       `json:"is_shared,omitempty"`
	Is_dislikes      bool       `json:"is_dislikes,omitempty"`
}

func (product *Eventos) TableName() string {
	return "eventos"
}
