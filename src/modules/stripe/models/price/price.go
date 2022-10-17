package price

type Price struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	Active            bool        `json:"active"`
	BillingScheme     string      `json:"billing_scheme"`
	Created           int64       `json:"created"`
	Currency          string      `json:"currency"`
	CustomUnitAmount  interface{} `json:"custom_unit_amount"`
	Livemode          bool        `json:"livemode"`
	LookupKey         interface{} `json:"lookup_key"`
	Metadata          Metadata    `json:"metadata"`
	Nickname          interface{} `json:"nickname"`
	Product           string      `json:"product"`
	Recurring         Recurring   `json:"recurring"`
	TaxBehavior       string      `json:"tax_behavior"`
	TiersMode         interface{} `json:"tiers_mode"`
	TransformQuantity interface{} `json:"transform_quantity"`
	Type              string      `json:"type"`
	UnitAmount        int64       `json:"unit_amount"`
	UnitAmountDecimal string      `json:"unit_amount_decimal"`
	Error             *Error      `json:"error"`
}

type Metadata struct {
}

type Recurring struct {
	AggregateUsage  interface{} `json:"aggregate_usage"`
	Interval        string      `json:"interval"`
	IntervalCount   int64       `json:"interval_count"`
	TrialPeriodDays interface{} `json:"trial_period_days"`
	UsageType       string      `json:"usage_type"`
}

type Error struct {
	Code    string `json:"code"`
	DocURL  string `json:"doc_url"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}
