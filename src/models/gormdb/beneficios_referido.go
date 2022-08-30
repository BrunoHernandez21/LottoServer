package gormdb

type Beneficios_referido struct {
	Activo       bool   `json:"activo"`
	Beneficio_id uint32 `json:"beneficio_id"`
}

func (product *Beneficios_plan) Beneficios_referido() string {
	return "beneficios_referido"
}
