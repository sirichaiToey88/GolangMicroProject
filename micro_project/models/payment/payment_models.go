package payment

type PaymentStripe struct {
	Amount      int64  `json:"amount"`
	Quantity    int64  `json:"quantity"`
	Description string `json:"description"`
	Name        string `json:"name"`
}
