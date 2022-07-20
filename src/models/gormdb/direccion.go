package gormdb

type Direccion struct {
	Id      uint32  `json:"id"`
	User_id *uint32 `json:"user_id,omitempty"`
	Tipo    *string `json:"tipo,omitempty"`
	Pais    *string `json:"pais,omitempty"`
	Ciudad  *string `json:"ciudad,omitempty"`
	Calle   *string `json:"calle,omitempty"`
	Cp      *string `json:"cp,omitempty"`
	Numero  *string `json:"numero,omitempty"`
}

func (product *Direccion) TableName() string {
	return "direccion"
}
