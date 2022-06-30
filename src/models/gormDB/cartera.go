package gormdb

type Carteras struct {
	Id                    uint32  `json:"id"`
	Activo                *bool   `json:"activo"`
	Acumulado_alto8am     *uint32 `json:"acumulado_alto8am,omitempty"`
	Acumulado_bajo8pm     *uint32 `json:"acumulado_bajo8pm,omitempty"`
	Aproximacion_alta00am *uint32 `json:"aproximacion_alta00am,omitempty"`
	Aproximacion_baja     *uint32 `json:"aproximacion_baja,omitempty"`
	Oportunidades         *uint32 `json:"iportunidades,omitempty"`
	Id_usuario            *uint32 `json:"id_usuario,omitempty"`
}

func (product *Carteras) TableName() string {
	return "carteras"
}
