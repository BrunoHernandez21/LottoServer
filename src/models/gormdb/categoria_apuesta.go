package gormdb

type Categoria_evento struct {
	Id     uint32  `json:"id"`
	Nombre *string `json:"nombre,omitempty"`
}

func (product *Categoria_evento) TableName() string {
	return "categoria_evento"
}
