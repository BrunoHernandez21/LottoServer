package product

type Product struct {
	ID                  string        `json:"id"`
	Object              string        `json:"object"`
	Active              bool          `json:"active"`
	Attributes          []interface{} `json:"attributes"`
	Created             int64         `json:"created"`
	DefaultPrice        interface{}   `json:"default_price"`
	Description         interface{}   `json:"description"`
	Images              []interface{} `json:"images"`
	Livemode            bool          `json:"livemode"`
	Metadata            Metadata      `json:"metadata"`
	Name                string        `json:"name"`
	PackageDimensions   interface{}   `json:"package_dimensions"`
	Shippable           interface{}   `json:"shippable"`
	StatementDescriptor interface{}   `json:"statement_descriptor"`
	TaxCode             interface{}   `json:"tax_code"`
	Type                string        `json:"type"`
	UnitLabel           interface{}   `json:"unit_label"`
	Updated             int64         `json:"updated"`
	URL                 interface{}   `json:"url"`
	Error               *Error        `json:"error"`
}

type Metadata struct {
	PlanID string `json:"plan_id"`
}

type Error struct {
	Code    string `json:"code"`
	DocURL  string `json:"doc_url"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}
