package payment

import (
	"fmt"
	"net/http"
	"project/models/payment"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"

	//"github.com/omise/omise-go/operations"

	"github.com/labstack/echo/v4"
)

func PaymentOmise(e echo.Context) error {
	body := &payment.PaymentStripe{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	// Check if any parameter in the body is null or empty
	if body.Amount == 0 {
		return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
	}

	// Set API Key
	publicKey := "pkey_test_5wn0zcp72wx26cq0s7l"
	secretKey := "skey_test_5wn0xwa83g3exm4q5el"
	omiseClient, err := omise.NewClient(publicKey, secretKey)
	if err != nil {
		fmt.Println("Error creating Omise client:", err)
		return e.JSON(http.StatusInternalServerError, "Error processing payment.")
	}

	// Create Recipient
	recipient, createRecipient := &omise.Recipient{}, &operations.CreateRecipient{
		Name:  "JOHN DOE",
		Email: "john.doe@example.com",
		Type:  "individual",
		BankAccount: &omise.BankAccount{
			Brand:  "bbl",
			Number: "1234567890",
			Name:   "JOHN DOE",
		},
	}

	if e := omiseClient.Do(recipient, createRecipient); e != nil {
		fmt.Println("Error creating recipient:", e)
		// return e.JSON(http.StatusInternalServerError, "Error processing payment.")
	}
	fmt.Println("recipient.ID", recipient.ID)
	// Create Transfer
	transfer, createTransfer := &omise.Transfer{}, &operations.CreateTransfer{
		Amount:    body.Amount,
		Recipient: recipient.ID,
	}

	if e := omiseClient.Do(transfer, createTransfer); e != nil {
		fmt.Println("Error creating transfer:", e)
		// return e.JSON(http.StatusInternalServerError, "Error processing payment.")
	}

	// Retrieve Transfer
	retrievedTransfer := &omise.Transfer{}
	if e := omiseClient.Do(retrievedTransfer, &operations.RetrieveTransfer{transfer.ID}); e != nil {
		fmt.Println("Error retrieving transfer:", e)
		// return e.JSON(http.StatusInternalServerError, "Error processing payment.")
	}

	fmt.Println("retrievedTransfer:", retrievedTransfer)
	fmt.Printf("Transfer ID: %s\n", retrievedTransfer.ID)
	fmt.Printf("Status: %s\n", retrievedTransfer.Paid)
	fmt.Printf("Amount: %d\n", retrievedTransfer.Amount)
	fmt.Printf("Currency: %s\n", retrievedTransfer.Currency)
	if retrievedTransfer.ID != "" {
		fmt.Println("Payment successful")
		return e.JSON(http.StatusOK, "Payment successful.")
	} else {
		fmt.Println("Payment failed")
		return e.JSON(http.StatusInternalServerError, "Payment failed.")
	}
}
