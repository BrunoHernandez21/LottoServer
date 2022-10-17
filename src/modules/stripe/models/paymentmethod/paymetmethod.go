package paymentmethod

type PaymentMetod struct {
	ID             string         `json:"id"`
	Object         string         `json:"object"`
	BillingDetails BillingDetails `json:"billing_details"`
	Card           Card           `json:"card"`
	Created        int64          `json:"created"`
	Customer       string         `json:"customer"`
	Livemode       bool           `json:"livemode"`
	Metadata       Metadata       `json:"metadata"`
	Type           string         `json:"type"`
	Error          *Error         `json:"error"`
}

type BillingDetails struct {
	Address Address `json:"address"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
}

type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
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

type Checks struct {
	AddressLine1Check      interface{} `json:"address_line1_check"`
	AddressPostalCodeCheck interface{} `json:"address_postal_code_check"`
	CvcCheck               string      `json:"cvc_check"`
}

type Networks struct {
	Available []string    `json:"available"`
	Preferred interface{} `json:"preferred"`
}

type ThreeDSecureUsage struct {
	Supported bool `json:"supported"`
}

type Metadata struct {
}

type Error struct {
	Code    string `json:"code"`
	DocURL  string `json:"doc_url"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}
