package impstripe

import (
	"encoding/json"
	"errors"
	"lottomusic/src/config"
	"lottomusic/src/modules/stripe/models/customer"
	"lottomusic/src/modules/stripe/models/paymentmethod"
	"lottomusic/src/modules/stripe/models/subscription"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Update_suscription_now(sub_id string, item_sub_id string, price_id string, orden_id uint32) (*subscription.Subscription, error) {
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("items[0][id]", item_sub_id)
	args.Set("items[0][price]", price_id)
	args.Set("items[0][metadata][orden_id]", strconv.FormatUint(uint64(orden_id), 10))
	args.Set("cancel_at_period_end", "false")
	args.Set("proration_behavior", "always_invoice")
	args.Set("metadata[orden_id]", strconv.FormatUint(uint64(orden_id), 10))
	a.Form(args)
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

func Update_suscription_proration(sub_id string, item_sub_id string, price_id string, orden_id uint32) (*subscription.Subscription, error) {
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("items[0][id]", item_sub_id)
	args.Set("items[0][price]", price_id)
	args.Set("cancel_at_period_end", "false")
	args.Set("proration_behavior", "create_prorations")
	args.Set("items[0][metadata][orden_id]", strconv.FormatUint(uint64(orden_id), 10))
	args.Set("metadata[orden_id]", strconv.FormatUint(uint64(orden_id), 10))
	a.Form(args)
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

func Update_customer(payment_id string, customer_id string) (*customer.Customer, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("invoice_settings[default_payment_method]", payment_id)
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/customers/" + customer_id)
	if err := a.Parse(); err != nil {
		return nil, err
	}
	_, intentBody, intentErr := a.Bytes()
	if intentErr != nil {
		return nil, nil
	}
	intent_out := customer.Customer{}
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

func Atach(customer string, payment_id string) (*paymentmethod.PaymentMetod, error) {
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("customer", customer)
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_methods/" + payment_id + "/attach")
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

func Detach(payment_id string) (*paymentmethod.PaymentMetod, error) {
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_methods/" + payment_id + "/detach")
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
