package gormdb

type Usuarios_roles struct {
	User_id uint32 `json:"user_id"`
	Role_id uint32 `json:"role_id"`
}

func (product *Usuarios_roles) TableName() string {
	return "Usuarios_roles"
}
