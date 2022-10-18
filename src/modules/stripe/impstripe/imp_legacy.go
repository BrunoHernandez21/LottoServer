package impstripe

import (
	"encoding/json"
	"errors"
	"lottomusic/src/config"
	"lottomusic/src/models/gormdb"
	"lottomusic/src/models/stripem"
	"lottomusic/src/modules/stripe/models/paymentintent"
	"lottomusic/src/modules/stripe/models/paymentmethod"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Payment(payM stripem.Stripe_Payment, orden *gormdb.Ordenes) (*paymentintent.PaymentIntent, error) {
	// Generate paymentMethod
	a := fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	args := fiber.AcquireArgs()
	args.Set("type", "card")
	args.Set("card[number]", payM.Card_number)
	args.Set("card[exp_month]", payM.Expiry_month)
	args.Set("card[cvc]", payM.Cvc)
	args.Set("card[exp_year]", payM.Expiry_year)
	a.Form(args)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer "+config.Stripekey)
	req.SetRequestURI("https://api.stripe.com/v1/payment_methods")
	if err := a.Parse(); err != nil {
		return nil, err
	}
	_, body, errs := a.Bytes()
	if errs != nil {
		ers := errors.New("error en la peticion a stripe payment_methods")
		return nil, ers
	}
	card_out := paymentmethod.PaymentMetod{}
	err := json.Unmarshal(body, &card_out)
	if err != nil {
		return nil, err
	}
	if card_out.Error != nil {
		outerr := errors.New(card_out.Error.Message)
		return nil, outerr
	}

	// Generate intent
	a = fiber.AcquireAgent()
	a.ContentType("application/x-www-form-urlencoded")
	// a.BasicAuth(config.Stripekey, "")
	args = fiber.AcquireArgs()
	args.Set("amount", strconv.FormatUint(uint64(orden.Precio_total*100), 10))
	args.Set("currency", "mxn")
	args.Set("payment_method", card_out.ID)
	args.Set("confirm", "true")
	args.Set("capture_method", "automatic")
	args.Set("metadata[orden]", strconv.FormatUint(uint64(orden.Id), 10))
	a.Form(args)
	req = a.Request()
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
	err = json.Unmarshal(intentBody, &intent_out)
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

/*
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("POST")
	a.ContentType("application/json")
	a.JSON(input)
	req.Header.Add("Authorization", "Basic "+globals.stripe)
	req.SetRequestURI("http://187.213.77.165:25565/api/compra/compra")
	if err := a.Parse(); err != nil {
		m["mensaje"] = err.Error()
		return c.Status(500).JSON(m)
	}
	code, body, errs := a.Bytes()*/
