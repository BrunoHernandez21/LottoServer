package customer

type Customer struct {
	ID                  string          `json:"id"`
	Object              string          `json:"object"`
	Address             interface{}     `json:"address"`
	Balance             int64           `json:"balance"`
	Created             int64           `json:"created"`
	Currency            interface{}     `json:"currency"`
	DefaultSource       interface{}     `json:"default_source"`
	Delinquent          bool            `json:"delinquent"`
	Description         interface{}     `json:"description"`
	Discount            interface{}     `json:"discount"`
	Email               interface{}     `json:"email"`
	InvoicePrefix       string          `json:"invoice_prefix"`
	InvoiceSettings     InvoiceSettings `json:"invoice_settings"`
	Livemode            bool            `json:"livemode"`
	Metadata            Metadata        `json:"metadata"`
	Name                string          `json:"name"`
	NextInvoiceSequence int64           `json:"next_invoice_sequence"`
	Phone               interface{}     `json:"phone"`
	PreferredLocales    []interface{}   `json:"preferred_locales"`
	Shipping            interface{}     `json:"shipping"`
	TaxExempt           string          `json:"tax_exempt"`
	TestClock           interface{}     `json:"test_clock"`
	Error               *Error          `json:"error"`
}

type InvoiceSettings struct {
	CustomFields         interface{} `json:"custom_fields"`
	DefaultPaymentMethod string      `json:"default_payment_method"`
	Footer               interface{} `json:"footer"`
	RenderingOptions     interface{} `json:"rendering_options"`
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
