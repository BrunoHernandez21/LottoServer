package gormdb

type Beneficios_plan struct {
	Id           uint32 `json:"id"`
	Activo       bool   `json:"activo"`
	Beneficio_id uint32 `json:"beneficio_id"`
	Plan_id      uint32 `json:"plan_id"`
}

func (product *Beneficios_plan) TableName() string {
	return "beneficios_plan"
}
