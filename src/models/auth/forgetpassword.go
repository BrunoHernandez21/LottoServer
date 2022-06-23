package auth

type Set_Forgetpassword struct {
	Mensaje string `json:"mensaje"`
}

type Get_forgetpassword struct {
	Email string `json:"email"`
}
