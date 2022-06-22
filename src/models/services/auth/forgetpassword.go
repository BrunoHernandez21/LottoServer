package auth

import gormdb "lottomusic/src/models/gormDB"

type Forgetpassword struct {
	Usuario *gormdb.Usuarios
	Mensaje string
}

type Auth_Get_forgetpassword struct {
	Usuario *gormdb.Usuarios
	Mensaje string
}
