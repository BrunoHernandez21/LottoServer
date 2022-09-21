package inputs

import "time"

type Set_login struct {
	User_id       uint32     `json:"user_id"`
	Access_token  string     `json:"access_token"`
	Token_type    *string    `gorm:"null" json:"token_type,omitempty"`
	Refresh_token *string    `gorm:"null" json:"refresh_token,omitempty"`
	Expires_in    *time.Time `gorm:"null" json:"expires_in,omitempty"`
	Scope         *string    `gorm:"null" json:"scope,omitempty"`
	Jti           *string    `gorm:"null" json:"jti,omitempty"`
}

type Get_Login struct {
	Email         *string `json:"email,omitempty"`
	Password      *string `json:"password,omitempty"`
	Grant_type    *string `json:"grant_type,omitempty"`
	Client_id     *string `json:"client_id,omitempty"`
	Client_secret *string `json:"client_secret,omitempty"`
}
type Get_signup struct {
	Email        *string `json:"email,omitempty"`
	Password     *string `json:"password,omitempty"`
	Referido_por *string `json:"referido_por,omitempty"`
}
type Get_ChangePassword struct {
	Password string `json:"password"`
}

type Get_forgetpassword struct {
	Email string `json:"email"`
}

type Get_token struct {
	Username      string
	Password      string
	Grant_type    string
	Client_id     string
	Client_secret string
}

type Get_User struct {
	Id              uint32
	Comentarios     uint32
	Fecha_creacion  string
	Likes           uint32
	Vistas          uint32
	Apuesta         uint32
	Tipo_apuesta_id uint32
	Usuario         uint32
}
