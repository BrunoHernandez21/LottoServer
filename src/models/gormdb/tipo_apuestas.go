package gormdb

type Tipo_apuesta struct {
	Id     uint32 `json:"id"`
	Nombre string `json:"nombre,omitempty"`
}

func (product *Tipo_apuesta) TableName() string {
	return "tipo_apuesta"
}
