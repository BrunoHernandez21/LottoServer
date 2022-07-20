package gormdb

type Referido struct {
	Id      uint32 `json:"id"`
	User_id uint32 `json:"user_id"`
	Codigo  string `json:"codigo,omitempty"`
	Cobrado bool   `json:"cobrado,omitempty"`
}

func (product *Referido) TableName() string {
	return "referido"
}
