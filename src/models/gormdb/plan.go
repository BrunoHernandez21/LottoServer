package gormdb

type Planes struct {
	Id             uint32   `json:"id"`
	Activo         *bool    `json:"activo"`
	Puntos         uint32   `json:"puntos"`
	Nombre         *string  `json:"nombre"`
	Precio         *float32 `json:"precio"`
	Moneda         string   `json:"moneda,omitempty"`
	Descuento_item *float32 `json:"descuento_item,omitempty"`
	Impuesto       *float32 `json:"impuesto,omitempty"`
	Precio_total   *float32 `json:"precio_total,omitempty"`
	Duracion_dias  int32    `json:"duracion_dias,omitempty"`
	Suscribcion    bool     `json:"suscribcion"`
}

func (product *Planes) TableName() string {
	return "planes"
}
