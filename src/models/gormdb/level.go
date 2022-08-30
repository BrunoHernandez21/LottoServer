package gormdb

type Level struct {
	Level string `json:"level"`
}

func (product *Level) TableName() string {
	return "level"
}
