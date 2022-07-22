package compuestas

import "time"

type CarritoPlan struct {
	Id                    uint32     `json:"id"`
	Activo                *bool      `json:"activo"`
	Status                *string    `json:"status"`
	Cantidad              int32      `json:"cantidad"`
	Total                 *float32   `json:"total"`
	Fecha_carrito         *time.Time `json:"fecha_carrito"`
	Usuario_id            uint32     `json:"usuario_id"`
	Acumulado_alto8am     *uint32    `json:"acumulado_alto8am,omitempty"`
	Acumulado_bajo8pm     *uint32    `json:"acumulado_bajo8pm,omitempty"`
	Aproximacion_alta00am *uint32    `json:"aproximacion_alta00am,omitempty"`
	Aproximacion_baja     *uint32    `json:"aproximacion_baja,omitempty"`
	Nombre                *string    `json:"nombre,omitempty"`
	Oportunidades         *uint32    `json:"oportunidades,omitempty"`
	Precio                *float32   `json:"precio,omitempty"`
	Suscribcion           bool       `json:"suscribcion,omitempty"`
	Pago_unico            bool       `json:"pago_unico,omitempty"`
}

func (product *CarritoPlan) TableName() string {
	return "carritoplan"
}
