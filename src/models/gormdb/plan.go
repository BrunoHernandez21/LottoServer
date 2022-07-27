package gormdb

type Planes struct {
	Id          uint32   `json:"id"`
	Activo      *bool    `json:"activo"`
	Cash        *uint32  `json:"cash"`
	Nombre      *string  `json:"nombre"`
	Precio      *float32 `json:"precio"`
	Moneda      string   `json:"moneda,omitempty"`
	Suscribcion bool     `json:"suscribcion"`
}

func (product *Planes) TableName() string {
	return "planes"
}
