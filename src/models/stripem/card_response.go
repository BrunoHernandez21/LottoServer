package stripem

type StripeCardResponse struct {
	ID             string         `json:"id"`
	Object         string         `json:"object"`
	BillingDetails BillingDetails `json:"billing_details"`
	Card           Card           `json:"card"`
	Created        int64          `json:"created"`
	Customer       interface{}    `json:"customer"`
	Livemode       bool           `json:"livemode"`
	Metadata       Metadata       `json:"metadata"`
	Type           string         `json:"type"`
	Error          *Error         `json:"error"`
}

type Networks struct {
	Available []string    `json:"available"`
	Preferred interface{} `json:"preferred"`
}

type ThreeDSecureUsage struct {
	Supported bool `json:"supported"`
}

type Error struct {
	Code    string `json:"code"`
	DocURL  string `json:"doc_url"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}

type Card struct {
	Brand             string            `json:"brand"`
	Checks            Checks            `json:"checks"`
	Country           string            `json:"country"`
	ExpMonth          int64             `json:"exp_month"`
	ExpYear           int64             `json:"exp_year"`
	Fingerprint       string            `json:"fingerprint"`
	Funding           string            `json:"funding"`
	GeneratedFrom     interface{}       `json:"generated_from"`
	Last4             string            `json:"last4"`
	Networks          Networks          `json:"networks"`
	ThreeDSecureUsage ThreeDSecureUsage `json:"three_d_secure_usage"`
	Wallet            interface{}       `json:"wallet"`
}
