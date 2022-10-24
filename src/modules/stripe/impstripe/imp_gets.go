package impstripe

import (
	"encoding/json"
	"errors"
	"lottomusic/src/config"
	"lottomusic/src/modules/stripe/models/paymentmethod"
	"lottomusic/src/modules/stripe/models/subscription"

	"github.com/gofiber/fiber/v2"
)

func Get_suscription(sub_id string) (*subscription.Subscription, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	req := a.Request()
	req.Header.SetMethod("GET")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/subscriptions/" + sub_id)
	if err := a.Parse(); err != nil {
		return nil, err
	}
	_, intentBody, intentErr := a.Bytes()
	if intentErr != nil {
		return nil, nil
	}
	intent_out := subscription.Subscription{}
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

func Get_paymet_method(payment string) (*paymentmethod.PaymentMetod, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	req := a.Request()
	req.Header.SetMethod("GET")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_methods/" + payment)
	if err := a.Parse(); err != nil {
		return nil, err
	}
	_, intentBody, intentErr := a.Bytes()
	if intentErr != nil {
		return nil, nil
	}
	intent_out := paymentmethod.PaymentMetod{}
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
