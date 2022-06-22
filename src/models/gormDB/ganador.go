package gormdb

type Gganador struct {
	Id         uint32  `json:"id"`
	Cantidad   *int32  `json:"cantidad,omitempty"`
	Concepto   *string `json:"concepto,omitempty"`
	Id_apuesta *uint32 `json:"id_apuesta,omitempty"`
	Id_ganador *uint32 `json:"id_ganador,omitempty"`
	Id_usuario *uint32 `json:"id_usuario,omitempty"`
}
