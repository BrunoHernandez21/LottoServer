package gormdb

type Usuarios_roles struct {
	Id      uint32 `json:"id"`
	User_id uint32 `json:"user_id"`
	Role_id uint32 `json:"role_id"`
}

func (product *Usuarios_roles) TableName() string {
	return "usuarios_roles"
}
