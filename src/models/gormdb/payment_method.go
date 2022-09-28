package gormdb

import (
	"lottomusic/src/models/stripem"
	"strconv"
)

type Payment_method struct {
	Id              uint32  `json:"id"`
	Activo          bool    `json:"activo"`
	Default_payment bool    `json:"default_payment"`
	Card_number     string  `json:"card_number"`
	Cvc             string  `json:"cvc"`
	Expiry_month    uint32  `json:"expiry_month"`
	Expiry_year     uint32  `json:"expiry_year"`
	Holder_name     string  `json:"holder_name"`
	Type            *string `json:"type"`
	Sub_type        *string `json:"sub_type"`
	Usuario_id      uint32  `json:"usuario_id"`
}

func (tarjeta *Payment_method) ToStripeMethod(cvc string) stripem.Stripe_Payment {
	year := strconv.FormatUint(uint64(tarjeta.Expiry_year), 10)
	month := strconv.FormatUint(uint64(tarjeta.Expiry_month), 10)
	return stripem.Stripe_Payment{
		Card_number:  tarjeta.Card_number,
		Expiry_month: month,
		Cvc:          cvc,
		Expiry_year:  year,
		Usuario_id:   tarjeta.Usuario_id,
		Holder_name:  tarjeta.Holder_name,
	}
}
func (product *Payment_method) TableName() string {
	return "payment_method"
}
