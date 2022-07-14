package gormdb

type Hibernate_sequence struct {
	Next_val *uint32 `json:"id"`
}

func (product *Hibernate_sequence) TableName() string {
	return "hibernate_sequence"
}
