package gormdb

type Carteras struct {
	Id         uint32 `json:"id"`
	Cash       uint32 `json:"cash"`
	Usuario_id uint32 `json:"usuario_id"`
}

func (product *Carteras) TableName() string {
	return "carteras"
}
