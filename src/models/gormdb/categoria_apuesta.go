package gormdb

type Categoria_apuesta struct {
	Id     uint32  `json:"id"`
	Nombre *string `json:"nombre,omitempty"`
}

func (product *Categoria_apuesta) TableName() string {
	return "Categoria_apuesta"
}
