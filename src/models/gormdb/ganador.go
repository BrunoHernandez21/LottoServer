package gormdb

type Ganador struct {
	Id                uint32  `json:"id"`
	Cantidad          *int32  `json:"cantidad,omitempty"`
	Concepto          *string `json:"concepto,omitempty"`
	Evento_id         *uint32 `json:"evento_id,omitempty"`
	Usuario_id        *uint32 `json:"usuario_id,omitempty"`
	Evento_usuario_id *uint32 `json:"evento_usuario_id,omitempty"`
}

func (product *Ganador) TableName() string {
	return "ganador"
}
