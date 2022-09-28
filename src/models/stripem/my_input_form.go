package stripem

type Stripe_Payment struct {
	Card_number  string `json:"card_number"`
	Cvc          string `json:"cvc"`
	Expiry_month string `json:"expiry_month"`
	Expiry_year  string `json:"expiry_year"`
	Holder_name  string `json:"holder_name"`
	Usuario_id   uint32 `json:"usuario_id"`
}
