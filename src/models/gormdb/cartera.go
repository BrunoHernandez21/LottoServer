package gormdb

type Carteras struct {
	Id         uint32 `json:"id"`
	Puntos     uint32 `json:"puntos"`
	Saldo_mxn  uint32 `json:"saldo_mxn"`
	Saldo_usd  uint32 `json:"saldo_usd"`
	Usuario_id uint32 `json:"usuario_id"`
}

func (product *Carteras) TableName() string {
	return "carteras"
}
