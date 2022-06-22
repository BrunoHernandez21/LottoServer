package auth

import "time"

type Put_login struct {
	Access_token  string     `json:"access_token"`
	Token_type    *string    `gorm:"null" json:"token_type,omitempty"`
	Refresh_token *string    `gorm:"null" json:"refresh_token,omitempty"`
	Expires_in    *time.Time `gorm:"null" json:"expires_in,omitempty"`
	Scope         *string    `gorm:"null" json:"scope,omitempty"`
	Jti           *string    `gorm:"null" json:"jti,omitempty"`
}

type Get_Login struct {
	Username      *string `json:"username,omitempty"`
	Password      *string `json:"password,omitempty"`
	Grant_type    *string `json:"grant_type,omitempty"`
	Client_id     *string `json:"client_id,omitempty"`
	Client_secret *string `json:"client_secret,omitempty"`
}
