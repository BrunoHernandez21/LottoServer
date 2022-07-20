package inputs

type Set_ChangePassword struct {
	Mensaje string `json:"mensaje"`
}

type Get_ChangePassword struct {
	Password string `json:"password"`
}
