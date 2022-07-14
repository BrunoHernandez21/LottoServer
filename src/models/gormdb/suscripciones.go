package gormdb

type Suscripciones struct {
	Id                    uint32  `json:"id"`
	Activo                *bool   `json:"activo"`
	Acumulado_alto8am     *uint32 `json:"acumulado_alto8am,omitempty"`
	Acumulado_bajo8pm     *uint32 `json:"acumulado_bajo8pm,omitempty"`
	Aproximacion_alta00am *uint32 `json:"aproximacion_alta00am,omitempty"`
	Aproximacion_baja     *uint32 `json:"aproximacion_baja,omitempty"`
	Fecha_inicio          *string `json:"fecha_inicio,omitempty"`
	Oportunidades         *uint32 `json:"iportunidades,omitempty"`
	Id_plan               *uint32 `json:"id_plan,omitempty"`
	Id_usuario            *uint32 `json:"id_usuario,omitempty"`
}

func (product *Suscripciones) TableName() string {
	return "Suscripciones"
}
