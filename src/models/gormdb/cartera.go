package gormdb

type Carteras struct {
	Id                    uint32 `json:"id"`
	Acumulado_alto8am     uint32 `json:"acumulado_alto8am"`
	Acumulado_bajo8pm     uint32 `json:"acumulado_bajo8pm"`
	Aproximacion_alta00am uint32 `json:"aproximacion_alta00am"`
	Aproximacion_baja     uint32 `json:"aproximacion_baja"`
	Oportunidades         uint32 `json:"oportunidades"`
	Usuario_id            uint32 `json:"usuario_id"`
}

func (product *Carteras) TableName() string {
	return "carteras"
}
