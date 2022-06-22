package gormdb

type Usuarios struct {
	Id               uint32 `json:"id"`
	Activo           bool   `json:"activo" gorm:"type:varbinary(1)"`
	Apellidom        string `json:"apellidom,omitempty"`
	Apellidop        string `json:"apellidop,omitempty"`
	Email            string `json:"email,omitempty"`
	Fecha_nacimiento string `json:"fecha_nacimiento,omitempty"`
	Nombre           string `json:"nombre,omitempty"`
	Password         string `json:"password,omitempty"`
	Telefono         string `json:"telefono,omitempty"`
}
