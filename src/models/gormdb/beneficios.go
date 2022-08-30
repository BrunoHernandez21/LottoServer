package gormdb

type Beneficios struct {
	Id       uint32  `json:"id"`
	Llave    string  `json:"llave"`
	Tipo     string  `json:"tipo"`
	Moneda   string  `json:"moneda"`
	Valor    float32 `json:"valor"`
	Dias     uint32  `json:"dias"`
	Acces_id string  `json:"acces_id"`
	Max_get  int32   `json:"max_get"`
}

func (product *Beneficios) TableName() string {
	return "beneficios"
}
