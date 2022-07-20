package gormdb

type Beneficios struct {
	Id              uint32  `json:"id"`
	Activo          bool    `json:"activo"`
	Llave           string  `json:"llave"`
	Tipo            string  `json:"tipo"`
	Moneda          string  `json:"moneda"`
	Valor           float32 `json:"valor"`
	Repetido        bool    `json:"repetido"`
	Suscripcion     bool    `json:"suscripcion"`
	Pago_individual bool    `json:"pago_individual"`
	Referido        bool    `json:"referido"`
}

func (product *Beneficios) TableName() string {
	return "beneficios"
}
