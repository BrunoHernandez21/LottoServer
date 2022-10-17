package paymentintent

type PaymentIntent struct {
	ID                        string               `json:"id"`
	Object                    string               `json:"object"`
	Amount                    int64                `json:"amount"`
	AmountCapturable          int64                `json:"amount_capturable"`
	AmountDetails             AmountDetails        `json:"amount_details"`
	AmountReceived            int64                `json:"amount_received"`
	Application               interface{}          `json:"application"`
	ApplicationFeeAmount      interface{}          `json:"application_fee_amount"`
	AutomaticPaymentMethods   interface{}          `json:"automatic_payment_methods"`
	CanceledAt                interface{}          `json:"canceled_at"`
	CancellationReason        interface{}          `json:"cancellation_reason"`
	CaptureMethod             string               `json:"capture_method"`
	Charges                   Charges              `json:"charges"`
	ClientSecret              string               `json:"client_secret"`
	ConfirmationMethod        string               `json:"confirmation_method"`
	Created                   int64                `json:"created"`
	Currency                  string               `json:"currency"`
	Customer                  interface{}          `json:"customer"`
	Description               string               `json:"description"`
	Invoice                   interface{}          `json:"invoice"`
	LastPaymentError          interface{}          `json:"last_payment_error"`
	Livemode                  bool                 `json:"livemode"`
	Metadata                  Metadata             `json:"metadata"`
	NextAction                interface{}          `json:"next_action"`
	OnBehalfOf                interface{}          `json:"on_behalf_of"`
	PaymentMethod             interface{}          `json:"payment_method"`
	PaymentMethodOptions      PaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes        []string             `json:"payment_method_types"`
	Processing                interface{}          `json:"processing"`
	ReceiptEmail              interface{}          `json:"receipt_email"`
	Review                    interface{}          `json:"review"`
	SetupFutureUsage          interface{}          `json:"setup_future_usage"`
	Shipping                  interface{}          `json:"shipping"`
	Source                    interface{}          `json:"source"`
	StatementDescriptor       interface{}          `json:"statement_descriptor"`
	StatementDescriptorSuffix interface{}          `json:"statement_descriptor_suffix"`
	Status                    string               `json:"status"`
	TransferData              interface{}          `json:"transfer_data"`
	TransferGroup             interface{}          `json:"transfer_group"`
	Error                     *Error               `json:"error"`
}

type AmountDetails struct {
	Tip Tip `json:"tip"`
}

type Tip struct {
}

type Charges struct {
	Object     string        `json:"object"`
	Data       []interface{} `json:"data"`
	HasMore    bool          `json:"has_more"`
	TotalCount int64         `json:"total_count"`
	URL        string        `json:"url"`
}

type Metadata struct {
	Orden string `json:"orden"`
}

type PaymentMethodOptions struct {
	Card Card `json:"card"`
}

type Card struct {
	Installments        interface{} `json:"installments"`
	MandateOptions      interface{} `json:"mandate_options"`
	Network             interface{} `json:"network"`
	RequestThreeDSecure string      `json:"request_three_d_secure"`
}

type Error struct {
	Code    string `json:"code"`
	DocURL  string `json:"doc_url"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}
