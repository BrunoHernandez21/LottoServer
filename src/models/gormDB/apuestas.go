package gormdb

type Apuestas struct {
	Id                   uint32   `json:"id"`
	Activo               *bool    `json:"activo"`
	Fechahoraapuesta     *string  `json:"fechahoraapuesta,omitempty"`
	Precio               *float32 `json:"precio,omitempty"`
	Acumulado            *float32 `json:"acumulado,omitempty"`
	Premio               *string  `json:"premio,omitempty"`
	Categoria_apuesta_id *uint32  `json:"categoria_apuesta_id,omitempty"`
	Tipo_apuesta_id      *uint32  `json:"tipo_apuesta_id,omitempty"`
	Video_id             *uint32  `json:"video_id,omitempty"`
}

func (product *Apuestas) TableName() string {
	return "Apuestas"
}
