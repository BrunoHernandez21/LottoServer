package stripem

type StripeIntentResponse struct {
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
	Description               interface{}          `json:"description"`
	Invoice                   interface{}          `json:"invoice"`
	LastPaymentError          interface{}          `json:"last_payment_error"`
	Livemode                  bool                 `json:"livemode"`
	Metadata                  Metadata             `json:"metadata"`
	NextAction                interface{}          `json:"next_action"`
	OnBehalfOf                interface{}          `json:"on_behalf_of"`
	PaymentMethod             string               `json:"payment_method"`
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

type Datum struct {
	ID                            string               `json:"id"`
	Object                        string               `json:"object"`
	Amount                        int64                `json:"amount"`
	AmountCaptured                int64                `json:"amount_captured"`
	AmountRefunded                int64                `json:"amount_refunded"`
	Application                   interface{}          `json:"application"`
	ApplicationFee                interface{}          `json:"application_fee"`
	ApplicationFeeAmount          interface{}          `json:"application_fee_amount"`
	BalanceTransaction            string               `json:"balance_transaction"`
	BillingDetails                BillingDetails       `json:"billing_details"`
	CalculatedStatementDescriptor string               `json:"calculated_statement_descriptor"`
	Captured                      bool                 `json:"captured"`
	Created                       int64                `json:"created"`
	Currency                      string               `json:"currency"`
	Customer                      interface{}          `json:"customer"`
	Description                   interface{}          `json:"description"`
	Destination                   interface{}          `json:"destination"`
	Dispute                       interface{}          `json:"dispute"`
	Disputed                      bool                 `json:"disputed"`
	FailureBalanceTransaction     interface{}          `json:"failure_balance_transaction"`
	FailureCode                   interface{}          `json:"failure_code"`
	FailureMessage                interface{}          `json:"failure_message"`
	FraudDetails                  Tip                  `json:"fraud_details"`
	Invoice                       interface{}          `json:"invoice"`
	Livemode                      bool                 `json:"livemode"`
	Metadata                      Metadata             `json:"metadata"`
	OnBehalfOf                    interface{}          `json:"on_behalf_of"`
	Order                         interface{}          `json:"order"`
	Outcome                       Outcome              `json:"outcome"`
	Paid                          bool                 `json:"paid"`
	PaymentIntent                 string               `json:"payment_intent"`
	PaymentMethod                 string               `json:"payment_method"`
	PaymentMethodDetails          PaymentMethodDetails `json:"payment_method_details"`
	ReceiptEmail                  interface{}          `json:"receipt_email"`
	ReceiptNumber                 interface{}          `json:"receipt_number"`
	ReceiptURL                    string               `json:"receipt_url"`
	Refunded                      bool                 `json:"refunded"`
	Refunds                       Charges              `json:"refunds"`
	Review                        interface{}          `json:"review"`
	Shipping                      interface{}          `json:"shipping"`
	Source                        interface{}          `json:"source"`
	SourceTransfer                interface{}          `json:"source_transfer"`
	StatementDescriptor           interface{}          `json:"statement_descriptor"`
	StatementDescriptorSuffix     interface{}          `json:"statement_descriptor_suffix"`
	Status                        string               `json:"status"`
	TransferData                  interface{}          `json:"transfer_data"`
	TransferGroup                 interface{}          `json:"transfer_group"`
}

type Charges struct {
	Object     string  `json:"object"`
	Data       []Datum `json:"data"`
	HasMore    bool    `json:"has_more"`
	TotalCount int64   `json:"total_count"`
	URL        string  `json:"url"`
}

type BillingDetails struct {
	Address Address     `json:"address"`
	Email   interface{} `json:"email"`
	Name    interface{} `json:"name"`
	Phone   interface{} `json:"phone"`
}

type Address struct {
	City       interface{} `json:"city"`
	Country    interface{} `json:"country"`
	Line1      interface{} `json:"line1"`
	Line2      interface{} `json:"line2"`
	PostalCode interface{} `json:"postal_code"`
	State      interface{} `json:"state"`
}

type Metadata struct {
	Orden string `json:"orden"`
}

type Outcome struct {
	NetworkStatus string      `json:"network_status"`
	Reason        interface{} `json:"reason"`
	RiskLevel     string      `json:"risk_level"`
	RiskScore     int64       `json:"risk_score"`
	SellerMessage string      `json:"seller_message"`
	Type          string      `json:"type"`
}

type PaymentMethodDetails struct {
	Card PaymentMethodDetailsCard `json:"card"`
	Type string                   `json:"type"`
}

type PaymentMethodDetailsCard struct {
	Brand        string      `json:"brand"`
	Checks       Checks      `json:"checks"`
	Country      string      `json:"country"`
	ExpMonth     int64       `json:"exp_month"`
	ExpYear      int64       `json:"exp_year"`
	Fingerprint  string      `json:"fingerprint"`
	Funding      string      `json:"funding"`
	Installments interface{} `json:"installments"`
	Last4        string      `json:"last4"`
	Mandate      interface{} `json:"mandate"`
	Network      string      `json:"network"`
	ThreeDSecure interface{} `json:"three_d_secure"`
	Wallet       interface{} `json:"wallet"`
}

type Checks struct {
	AddressLine1Check      interface{} `json:"address_line1_check"`
	AddressPostalCodeCheck interface{} `json:"address_postal_code_check"`
	CvcCheck               string      `json:"cvc_check"`
}

type PaymentMethodOptions struct {
	Card PaymentMethodOptionsCard `json:"card"`
}

type PaymentMethodOptionsCard struct {
	Installments        interface{} `json:"installments"`
	MandateOptions      interface{} `json:"mandate_options"`
	Network             interface{} `json:"network"`
	RequestThreeDSecure string      `json:"request_three_d_secure"`
}
