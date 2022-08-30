package gormdb

type Acces struct {
	Id           uint32 `json:"id"`
	Activo       bool   `json:"activo"`
	Beneficio_id uint32 `json:"beneficio_id"`
	Plan_id      uint32 `json:"plan_id"`
}

func (product *Acces) TableName() string {
	return "acces"
}
