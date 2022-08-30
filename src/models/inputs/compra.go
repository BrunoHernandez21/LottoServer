package inputs

type Checkout struct {
	Card_id uint32 `json:"card_id"`
}

type Get_Stripe struct {
	Orden_id  uint32 `json:"orden_id"`
	StripeKey string `json:"stripe_key"`
}
