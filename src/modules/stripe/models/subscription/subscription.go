package subscription

type Subscription struct {
	ID                            string          `json:"id"`
	Object                        string          `json:"object"`
	Application                   interface{}     `json:"application"`
	ApplicationFeePercent         interface{}     `json:"application_fee_percent"`
	AutomaticTax                  AutomaticTax    `json:"automatic_tax"`
	BillingCycleAnchor            int64           `json:"billing_cycle_anchor"`
	BillingThresholds             interface{}     `json:"billing_thresholds"`
	CancelAt                      interface{}     `json:"cancel_at"`
	CancelAtPeriodEnd             bool            `json:"cancel_at_period_end"`
	CanceledAt                    interface{}     `json:"canceled_at"`
	CollectionMethod              string          `json:"collection_method"`
	Created                       int64           `json:"created"`
	Currency                      string          `json:"currency"`
	CurrentPeriodEnd              int64           `json:"current_period_end"`
	CurrentPeriodStart            int64           `json:"current_period_start"`
	Customer                      string          `json:"customer"`
	DaysUntilDue                  interface{}     `json:"days_until_due"`
	DefaultPaymentMethod          interface{}     `json:"default_payment_method"`
	DefaultSource                 interface{}     `json:"default_source"`
	DefaultTaxRates               []interface{}   `json:"default_tax_rates"`
	Description                   interface{}     `json:"description"`
	Discount                      interface{}     `json:"discount"`
	EndedAt                       interface{}     `json:"ended_at"`
	Items                         Items           `json:"items"`
	LatestInvoice                 string          `json:"latest_invoice"`
	Livemode                      bool            `json:"livemode"`
	Metadata                      DatumMetadata   `json:"metadata"`
	NextPendingInvoiceItemInvoice interface{}     `json:"next_pending_invoice_item_invoice"`
	PauseCollection               interface{}     `json:"pause_collection"`
	PaymentSettings               PaymentSettings `json:"payment_settings"`
	PendingInvoiceItemInterval    interface{}     `json:"pending_invoice_item_interval"`
	PendingSetupIntent            interface{}     `json:"pending_setup_intent"`
	PendingUpdate                 interface{}     `json:"pending_update"`
	Plan                          Plan            `json:"plan"`
	Quantity                      int64           `json:"quantity"`
	Schedule                      interface{}     `json:"schedule"`
	StartDate                     int64           `json:"start_date"`
	Status                        string          `json:"status"`
	TestClock                     interface{}     `json:"test_clock"`
	TransferData                  interface{}     `json:"transfer_data"`
	TrialEnd                      interface{}     `json:"trial_end"`
	TrialStart                    interface{}     `json:"trial_start"`
	Error                         *Error          `json:"error"`
}

type AutomaticTax struct {
	Enabled bool `json:"enabled"`
}

type Items struct {
	Object     string  `json:"object"`
	Data       []Datum `json:"data"`
	HasMore    bool    `json:"has_more"`
	TotalCount int64   `json:"total_count"`
	URL        string  `json:"url"`
}

type Datum struct {
	ID                string        `json:"id"`
	Object            string        `json:"object"`
	BillingThresholds interface{}   `json:"billing_thresholds"`
	Created           int64         `json:"created"`
	Metadata          DatumMetadata `json:"metadata"`
	Plan              Plan          `json:"plan"`
	Price             Price         `json:"price"`
	Quantity          int64         `json:"quantity"`
	Subscription      string        `json:"subscription"`
	TaxRates          []interface{} `json:"tax_rates"`
}

type DatumMetadata struct {
	OrdenID string `json:"orden_id"`
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
	OrdenID string `json:"orden_id"`
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

type PaymentSettings struct {
	PaymentMethodOptions     interface{} `json:"payment_method_options"`
	PaymentMethodTypes       interface{} `json:"payment_method_types"`
	SaveDefaultPaymentMethod string      `json:"save_default_payment_method"`
}

type Error struct {
	Code    string `json:"code"`
	DocURL  string `json:"doc_url"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}
