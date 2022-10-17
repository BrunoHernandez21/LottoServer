package gormdb

type Planes struct {
	Id             uint32   `json:"id"`
	Activo         *bool    `json:"activo"`
	Titulo         *string  `json:"titulo"`
	Descripcion    *string  `json:"descripcion"`
	Pre_puntos     uint32   `json:"Pre_puntos"`
	Puntos         uint32   `json:"puntos"`
	Pre_precio     *float32 `json:"pre_precio,omitempty"`
	Precio         *float32 `json:"precio,omitempty"`
	Moneda         string   `json:"moneda,omitempty"`
	Suscribcion    bool     `json:"suscribcion"`
	Stripe_price   *string  `json:"stripe_price"`
	Stripe_product *string  `json:"stripe_product"`
}

func (product *Planes) TableName() string {
	return "planes"
}
