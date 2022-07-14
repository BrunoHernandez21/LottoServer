package gormdb

type Pago struct {
	Id              uint32  `json:"id"`
	Card_number     *string `json:"card_number"`
	Cvc             uint32  `json:"cvc"`
	Default_payment int8    `json:"default_payment"`
	Expiry_month    uint32  `json:"expiry_month"`
	Expiry_year     uint32  `json:"expiry_year"`
	Holder_name     *string `json:"holder_name"`
	Type            *string `json:"type"`
}

func (product *Pago) TableName() string {
	return "pago"
}
