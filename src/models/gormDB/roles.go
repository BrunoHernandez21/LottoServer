package gormdb

type Roles struct {
	Id     uint32 `json:"id"`
	Nombre string `json:"nombre,omitempty"`
}
