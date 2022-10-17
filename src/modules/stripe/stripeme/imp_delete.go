package stripeme

import (
	"encoding/json"
	"errors"
	"lottomusic/src/config"
	"lottomusic/src/modules/stripe/models/paymentintent"

	"github.com/gofiber/fiber/v2"
)

func Delete_suscription(sub_id string) (*paymentintent.PaymentIntent, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/subscriptions/" + sub_id)
	if err := a.Parse(); err != nil {
		return nil, err
	}
	_, intentBody, intentErr := a.Bytes()
	if intentErr != nil {
		return nil, nil
	}
	intent_out := paymentintent.PaymentIntent{}
	err := json.Unmarshal(intentBody, &intent_out)
	if err != nil {
		return nil, err
	}
	if intent_out.Error != nil {
		outerr := errors.New(intent_out.Error.Message)
		return nil, outerr
	}
	//out
	return &intent_out, nil
}
