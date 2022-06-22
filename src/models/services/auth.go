package services

import gormdb "lottomusic/src/models/gormDB"

type Auth_login struct {
	Access_token  string  `json:"access_token"`
	Token_type    *string `gorm:"null" json:"token_type,omitempty"`
	Refresh_token *string `gorm:"null" json:"refresh_token,omitempty"`
	Expires_in    *uint32 `gorm:"null" json:"expires_in,omitempty"`
	Scope         *string `gorm:"null" json:"scope,omitempty"`
	Jti           *string `gorm:"null" json:"jti,omitempty"`
}

type Auth_signup struct {
	Usuario string
	Mensaje string
}

type Auth_forgetpassword struct {
	Usuario *gormdb.Usuarios
	Mensaje string
}

type Auth_myuser struct {
	Usuario string
	Mensaje string
}
type Auth_delete_user struct {
	Mensaje string
}
type Auth_updateusuario struct {
	Usuario string
	Mensaje string
}

///////////////////////////////////
/// get
type Auth_Get_Login struct {
	Username      *string `json:"username,omitempty"`
	Password      *string `json:"password,omitempty"`
	Grant_type    *string `json:"grant_type,omitempty"`
	Client_id     *string `json:"client_id,omitempty"`
	Client_secret *string `json:"client_secret,omitempty"`
}

type Auth_Get_Users struct {
	Users []*gormdb.Usuarios
}

type Auth_Get_signup struct {
	Usuario  string
	Password string
}

type Auth_Get_token struct {
	Username      string
	Password      string
	Grant_type    string
	Client_id     string
	Client_secret string
}

type Auth_Get_User struct {
	Id              uint32
	Comentarios     uint32
	Fecha_creacion  string
	Likes           uint32
	Vistas          uint32
	Apuesta         uint32
	Tipo_apuesta_id uint32
	Usuario         uint32
}

type Auth_Get_forgetpassword struct {
	Usuario *gormdb.Usuarios
	Mensaje string
}

type Auth_Get_updateusuario struct {
	Usuario string
}
