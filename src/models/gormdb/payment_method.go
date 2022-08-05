package gormdb

type Payment_method struct {
	Id              uint32  `json:"id"`
	Activo          bool    `json:"activo"`
	Default_payment bool    `json:"default_payment"`
	Card_number     *string `json:"card_number"`
	Cvc             uint32  `json:"cvc"`
	Expiry_month    uint32  `json:"expiry_month"`
	Expiry_year     uint32  `json:"expiry_year"`
	Holder_name     *string `json:"holder_name"`
	Type            *string `json:"type"`
	Usuario_id      uint32  `json:"usuario_id"`
}

func (product *Payment_method) TableName() string {
	return "payment_method"
}
