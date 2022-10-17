package invoice

type Invoice struct {
	ID                           string            `json:"id"`
	Object                       string            `json:"object"`
	AccountCountry               string            `json:"account_country"`
	AccountName                  string            `json:"account_name"`
	AccountTaxIDS                interface{}       `json:"account_tax_ids"`
	AmountDue                    int64             `json:"amount_due"`
	AmountPaid                   int64             `json:"amount_paid"`
	AmountRemaining              int64             `json:"amount_remaining"`
	Application                  interface{}       `json:"application"`
	ApplicationFeeAmount         interface{}       `json:"application_fee_amount"`
	AttemptCount                 int64             `json:"attempt_count"`
	Attempted                    bool              `json:"attempted"`
	AutoAdvance                  bool              `json:"auto_advance"`
	AutomaticTax                 AutomaticTax      `json:"automatic_tax"`
	BillingReason                string            `json:"billing_reason"`
	Charge                       string            `json:"charge"`
	CollectionMethod             string            `json:"collection_method"`
	Created                      int64             `json:"created"`
	Currency                     string            `json:"currency"`
	CustomFields                 interface{}       `json:"custom_fields"`
	Customer                     string            `json:"customer"`
	CustomerAddress              interface{}       `json:"customer_address"`
	CustomerEmail                interface{}       `json:"customer_email"`
	CustomerName                 string            `json:"customer_name"`
	CustomerPhone                interface{}       `json:"customer_phone"`
	CustomerShipping             interface{}       `json:"customer_shipping"`
	CustomerTaxExempt            string            `json:"customer_tax_exempt"`
	CustomerTaxIDS               []interface{}     `json:"customer_tax_ids"`
	DefaultPaymentMethod         interface{}       `json:"default_payment_method"`
	DefaultSource                interface{}       `json:"default_source"`
	DefaultTaxRates              []interface{}     `json:"default_tax_rates"`
	Description                  interface{}       `json:"description"`
	Discount                     interface{}       `json:"discount"`
	Discounts                    []interface{}     `json:"discounts"`
	DueDate                      interface{}       `json:"due_date"`
	EndingBalance                int64             `json:"ending_balance"`
	Footer                       interface{}       `json:"footer"`
	FromInvoice                  interface{}       `json:"from_invoice"`
	HostedInvoiceURL             string            `json:"hosted_invoice_url"`
	InvoicePDF                   string            `json:"invoice_pdf"`
	LastFinalizationError        interface{}       `json:"last_finalization_error"`
	LatestRevision               interface{}       `json:"latest_revision"`
	Lines                        Lines             `json:"lines"`
	Livemode                     bool              `json:"livemode"`
	Metadata                     PlanMetadata      `json:"metadata"`
	NextPaymentAttempt           interface{}       `json:"next_payment_attempt"`
	Number                       string            `json:"number"`
	OnBehalfOf                   interface{}       `json:"on_behalf_of"`
	Paid                         bool              `json:"paid"`
	PaidOutOfBand                bool              `json:"paid_out_of_band"`
	PaymentIntent                string            `json:"payment_intent"`
	PaymentSettings              PaymentSettings   `json:"payment_settings"`
	PeriodEnd                    int64             `json:"period_end"`
	PeriodStart                  int64             `json:"period_start"`
	PostPaymentCreditNotesAmount int64             `json:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount  int64             `json:"pre_payment_credit_notes_amount"`
	Quote                        interface{}       `json:"quote"`
	ReceiptNumber                interface{}       `json:"receipt_number"`
	RenderingOptions             interface{}       `json:"rendering_options"`
	StartingBalance              int64             `json:"starting_balance"`
	StatementDescriptor          interface{}       `json:"statement_descriptor"`
	Status                       string            `json:"status"`
	StatusTransitions            StatusTransitions `json:"status_transitions"`
	Subscription                 string            `json:"subscription"`
	Subtotal                     int64             `json:"subtotal"`
	SubtotalExcludingTax         int64             `json:"subtotal_excluding_tax"`
	Tax                          interface{}       `json:"tax"`
	TestClock                    interface{}       `json:"test_clock"`
	Total                        int64             `json:"total"`
	TotalDiscountAmounts         []interface{}     `json:"total_discount_amounts"`
	TotalExcludingTax            int64             `json:"total_excluding_tax"`
	TotalTaxAmounts              []interface{}     `json:"total_tax_amounts"`
	TransferData                 interface{}       `json:"transfer_data"`
	WebhooksDeliveredAt          int64             `json:"webhooks_delivered_at"`
	Error                        *Error            `json:"error"`
}

type AutomaticTax struct {
	Enabled bool        `json:"enabled"`
	Status  interface{} `json:"status"`
}

type Lines struct {
	Object     string  `json:"object"`
	Data       []Datum `json:"data"`
	HasMore    bool    `json:"has_more"`
	TotalCount int64   `json:"total_count"`
	URL        string  `json:"url"`
}

type Datum struct {
	ID                     string           `json:"id"`
	Object                 string           `json:"object"`
	Amount                 int64            `json:"amount"`
	AmountExcludingTax     int64            `json:"amount_excluding_tax"`
	Currency               string           `json:"currency"`
	Description            string           `json:"description"`
	DiscountAmounts        []interface{}    `json:"discount_amounts"`
	Discountable           bool             `json:"discountable"`
	Discounts              []interface{}    `json:"discounts"`
	Livemode               bool             `json:"livemode"`
	Metadata               DatumMetadata    `json:"metadata"`
	Period                 Period           `json:"period"`
	Plan                   Plan             `json:"plan"`
	Price                  Price            `json:"price"`
	Proration              bool             `json:"proration"`
	ProrationDetails       ProrationDetails `json:"proration_details"`
	Quantity               int64            `json:"quantity"`
	Subscription           string           `json:"subscription"`
	SubscriptionItem       string           `json:"subscription_item"`
	TaxAmounts             []interface{}    `json:"tax_amounts"`
	TaxRates               []interface{}    `json:"tax_rates"`
	Type                   string           `json:"type"`
	UnitAmountExcludingTax string           `json:"unit_amount_excluding_tax"`
}

type DatumMetadata struct {
	OrdenID string `json:"orden_id"`
}

type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

type Plan struct {
	ID              string       `json:"id"`
	Object          string       `json:"object"`
	Active          bool         `json:"active"`
	AggregateUsage  interface{}  `json:"aggregate_usage"`
	Amount          int64        `json:"amount"`
	AmountDecimal   string       `json:"amount_decimal"`
	BillingScheme   string       `json:"billing_scheme"`
	Created         int64        `json:"created"`
	Currency        string       `json:"currency"`
	Interval        string       `json:"interval"`
	IntervalCount   int64        `json:"interval_count"`
	Livemode        bool         `json:"livemode"`
	Metadata        PlanMetadata `json:"metadata"`
	Nickname        interface{}  `json:"nickname"`
	Product         string       `json:"product"`
	TiersMode       interface{}  `json:"tiers_mode"`
	TransformUsage  interface{}  `json:"transform_usage"`
	TrialPeriodDays interface{}  `json:"trial_period_days"`
	UsageType       string       `json:"usage_type"`
}

type PlanMetadata struct {
}

type Price struct {
	ID                string       `json:"id"`
	Object            string       `json:"object"`
	Active            bool         `json:"active"`
	BillingScheme     string       `json:"billing_scheme"`
	Created           int64        `json:"created"`
	Currency          string       `json:"currency"`
	CustomUnitAmount  interface{}  `json:"custom_unit_amount"`
	Livemode          bool         `json:"livemode"`
	LookupKey         interface{}  `json:"lookup_key"`
	Metadata          PlanMetadata `json:"metadata"`
	Nickname          interface{}  `json:"nickname"`
	Product           string       `json:"product"`
	Recurring         Recurring    `json:"recurring"`
	TaxBehavior       string       `json:"tax_behavior"`
	TiersMode         interface{}  `json:"tiers_mode"`
	TransformQuantity interface{}  `json:"transform_quantity"`
	Type              string       `json:"type"`
	UnitAmount        int64        `json:"unit_amount"`
	UnitAmountDecimal string       `json:"unit_amount_decimal"`
}

type Recurring struct {
	AggregateUsage  interface{} `json:"aggregate_usage"`
	Interval        string      `json:"interval"`
	IntervalCount   int64       `json:"interval_count"`
	TrialPeriodDays interface{} `json:"trial_period_days"`
	UsageType       string      `json:"usage_type"`
}

type ProrationDetails struct {
	CreditedItems interface{} `json:"credited_items"`
}

type PaymentSettings struct {
	DefaultMandate       interface{} `json:"default_mandate"`
	PaymentMethodOptions interface{} `json:"payment_method_options"`
	PaymentMethodTypes   interface{} `json:"payment_method_types"`
}

type StatusTransitions struct {
	FinalizedAt           int64       `json:"finalized_at"`
	MarkedUncollectibleAt interface{} `json:"marked_uncollectible_at"`
	PaidAt                int64       `json:"paid_at"`
	VoidedAt              interface{} `json:"voided_at"`
}

type Error struct {
	Code    string `json:"code"`
	DocURL  string `json:"doc_url"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}
