package inputs

type Checkout struct {
	IDs []uint32 `json:"IDs"`
}

type Get_Stripe struct {
	Orden_id  uint32 `json:"orden_id"`
	StripeKey string `json:"stripe_key"`
}
