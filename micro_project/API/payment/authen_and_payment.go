package payment

import (
	"fmt"
	"net/http"
	"project/models/payment"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func AuthenAndPayment(e echo.Context) error {

	body := &payment.PaymentStripe{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	//Check if any parameter in the body is null or empty
	if body.Quantity == 0 || body.Name == "" || body.Amount == 0 {
		return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
	}
	// if body.Amount == 0 {
	// 	body.Amount = 1000 
	// }
	// if body.Quantity == 0 {
	// 	body.Quantity = 1000 
	// }
	fmt.Println("body.Amount", body.Amount)
	stripe.Key = "sk_test_51NabEdJdRkwDA65xFJQ01gYvCYGt0Qu7DRwAeQ7zGnO3yZgDhjH4fjChqZg5heF5SOQ8ZhyCP1Vg9eYiEspg1tbq00YaaeYoYg"
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://your-website.com/success"),
		CancelURL:  stripe.String("https://your-website.com/cancel"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		// Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Amount:      stripe.Int64(body.Amount),
				Currency:    stripe.String("THB"),
				Quantity:    stripe.Int64(body.Quantity),
				Description: stripe.String(body.Description),
				Name:        stripe.String(body.Name),
			},
		},
	}

	session, err := session.New(params)
	if err != nil {
		fmt.Println("err =>> :", err)
	}

	fmt.Println("Checkout URL:", session.SuccessURL)
	return nil

}
