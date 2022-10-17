package stripeme

import (
	"encoding/json"
	"errors"
	"lottomusic/src/config"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/stripem"
	"lottomusic/src/modules/stripe/models/customer"
	"lottomusic/src/modules/stripe/models/paymentintent"
	"lottomusic/src/modules/stripe/models/paymentmethod"
	"lottomusic/src/modules/stripe/models/price"
	"lottomusic/src/modules/stripe/models/product"
	"lottomusic/src/modules/stripe/models/subscription"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Create_payment_intent(orden *gormdb.Ordenes) (*paymentintent.PaymentIntent, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("amount", strconv.FormatUint(uint64(orden.Precio_total*100), 10))
	args.Set("currency", "mxn")
	args.Set("description", "Lotto Music points")
	args.Set("metadata[orden]", strconv.FormatUint(uint64(orden.Id), 10))
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_intents")
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

func Pay_payment_intent_id(orden *gormdb.Ordenes, payment_id string, ordenid string) (*paymentintent.PaymentIntent, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("amount", strconv.FormatUint(uint64(orden.Precio_total*100), 10))
	args.Set("currency", "mxn")
	args.Set("description", "Lotto Music points")
	args.Set("metadata[orden]", strconv.FormatUint(uint64(orden.Id), 10))
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_intents")
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

func Pay_payment_intent(orden *gormdb.Ordenes, Paymentid string) (*paymentintent.PaymentIntent, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("amount", strconv.FormatUint(uint64(orden.Precio_total*100), 10))
	args.Set("currency", "mxn")
	args.Set("payment_method", Paymentid)
	args.Set("description", "Lotto Music points")
	args.Set("metadata[orden]", strconv.FormatUint(uint64(orden.Id), 10))
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_intents")
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

func Create_payment_method(orden *stripem.Stripe_Payment) (*paymentmethod.PaymentMetod, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	// a.BasicAuth(config.Stripekey, "")
	args := fiber.AcquireArgs()
	args.Set("card[number]", orden.Card_number)
	args.Set("card[exp_year]", orden.Expiry_year)
	args.Set("card[exp_month]", orden.Expiry_month)
	args.Set("card[cvc]", orden.Cvc)
	args.Set("type", "card")
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_methods")
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

func Create_product(plan *gormdb.Planes) (*product.Product, error) {
	// Generate intent

	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("name", *plan.Titulo)
	args.Set("metadata[plan_id]", strconv.FormatUint(uint64(plan.Id), 10))
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/products")
	if err := a.Parse(); err != nil {
		return nil, err
	}
	_, intentBody, intentErr := a.Bytes()
	if intentErr != nil {
		return nil, nil
	}
	intent_out := product.Product{}
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

func Create_price(plan *gormdb.Planes, productid string) (*price.Price, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("unit_amount", strconv.FormatUint(uint64(*plan.Precio*100), 10))
	args.Set("currency", "mxn")
	args.Set("recurring[interval]", "month")
	args.Set("product", productid)
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/prices")
	if err := a.Parse(); err != nil {
		return nil, err
	}
	_, intentBody, intentErr := a.Bytes()
	if intentErr != nil {
		return nil, nil
	}
	intent_out := price.Price{}
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

func Create_customer(payment_id string, user_id string) (*customer.Customer, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("payment_method", payment_id)
	args.Set("name", user_id)
	args.Set("invoice_settings[default_payment_method]", payment_id)
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/customers")
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

func Create_suscription(plan gormdb.Planes, customer string) (*subscription.Subscription, error) {
	// Generate intent
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("customer", customer)
	args.Set("items[0][price]", *plan.Stripe_price)
	args.Set("items[0][metadata][orden_id]", strconv.FormatUint(uint64(plan.Id), 10))
	args.Set("metadata[orden_id]", strconv.FormatUint(uint64(plan.Id), 10))
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/subscriptions")
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
