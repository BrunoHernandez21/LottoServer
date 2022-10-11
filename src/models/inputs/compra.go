package inputs

type Checkout struct {
	Card_id uint32 `json:"card_id"`
	Cvc     string `json:"cvc"`
}

type RetryCheckout struct {
	Card_id  uint32 `json:"card_id"`
	Orden_id uint32 `json:"orden_id"`
	Cvc      string `json:"cvc"`
}
type GenerarPaymentItent struct {
	Orden_id uint32 `json:"orden_id"`
}

type GenerarOrden struct {
	Orden_id uint32 `json:"card_id"`
	Cvc      string `json:"cvc"`
}
