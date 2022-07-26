package views

type Plan_Suscribcion struct {
	Id                    uint32   `json:"id"`
	Activo                *bool    `json:"activo"`
	Acumulado_alto8am     *float32 `json:"status"`
	Acumulado_bajo8pm     *float32 `json:"acumulado_alto8am,omitempty"`
	Aproximacion_alta00am *float32 `json:"acumulado_bajo8pm"`
	Aproximacion_baja     *float32 `json:"aproximacion_alta00am"`
	Oportunidades         *float32 `json:"aproximacion_baja"`
	Nombre                *string  `json:"oportunidades"`
	Precio                *float32 `json:"nombre"`
	Suscribcion           *string  `json:"precio,omitempty"`
}

func (product *EventoVideo) Plan_Suscribcion() string {
	return "eventos_videos"
}
