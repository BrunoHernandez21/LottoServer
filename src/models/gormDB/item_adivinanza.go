package gormdb

type Item_adivinanza struct {
	Id                  uint32  `json:"id"`
	Cantidad            *int32  `json:"cantidad,omitempty"`
	Comentarios_usuario float32 `json:"comentarios_usuario"`
	Id_apuesta          *uint32 `json:"id_apuesta,omitempty"`
	Juego_id            *uint32 `json:"juego_id,omitempty"`
}
