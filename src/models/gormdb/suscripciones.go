package gormdb

import "time"

type Suscripciones struct {
	Monto_mensual      *float32   `json:"monto_mensual"`
	Fecha_inicio       *time.Time `json:"fecha_inicio"`
	Fecha_fin          *time.Time `json:"fecha_fin"`
	Plan_id            uint32     `json:"plan_id"`
	Usuario_id         uint32     `json:"usuario_id"`
	Stripe_customer    string     `json:"stripe_customer"`
	Stripe_suscription string     `json:"stripe_suscription"`
	Stripe_paymenth    string     `json:"stripe_paymenth"`
}

func (product *Suscripciones) TableName() string {
	return "suscripciones"
}
