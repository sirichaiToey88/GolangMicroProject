package shoppingCart

import (
	"fmt"
	"log"
	"net/http"
	shoppingcart "project/models/shopping_cart"
	"project/schemas"

	"github.com/labstack/echo/v4"
)

type DataCart struct {
	User_id string `json:"user_id"`
}

func SearchOrdeyByID(e echo.Context) error {
	body := &DataCart{}
	dataCart := []shoppingcart.ShoppingSearch{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	db, err := schemas.Connectdb()

	// rows, err := db.Query(`SELECT shopping_cart.id, shopping_cart.user_id,shopping_cart.product_id,shopping_cart.product_title,shopping_cart.Product_price,shopping_cart.total_amount,shopping_cart.quantity_product,shopping_cart.payment_type,shopping_cart.order_id,shopping_cart.create_date,shopping_cart.modify_date,shopping_cart.image_url,payment.id,payment.order_id,payment.account,payment.source,payment.distination,payment.total,payment.create_date,payment.modify_date FROM shopping_cart  INNER JOIN payment ON payment.order_id = shopping_cart.order_id WHERE user_id = ?`, body.User_id)
	rows, err := db.Query(`SELECT *
	FROM shopping_cart sc
	INNER JOIN payment py ON sc.order_id = py.order_id
	INNER JOIN main_cart mc ON sc.order_id = mc.order_id
	WHERE sc.user_id = ? ORDER BY sc.create_date DESC`, body.User_id)

	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, rows)
	}
	defer rows.Close()

	rowsSum, err := db.Query(`SELECT sum(total_amount) FROM shopping_cart WHERE user_id = ?`, body.User_id)

	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, rowsSum)
	}
	defer rowsSum.Close()

	total := 0.0
	for rows.Next() {
		var rawCart shoppingcart.ShoppingSearch
		var rawMainCart shoppingcart.MainCart

		if err = rows.Scan(&rawCart.Id, &rawCart.User_id, &rawCart.Product_id, &rawCart.Product_title, &rawCart.Product_price, &rawCart.Total_amount,
			&rawCart.Quantity_product, &rawCart.Payment_type, &rawCart.Order_id, &rawCart.Image_url, &rawCart.Create_date, &rawCart.Modify_date,
			&rawCart.Payment.Id, &rawCart.Payment.Order_id, &rawCart.Payment.Account,
			&rawCart.Payment.Source, &rawCart.Payment.Distination, &rawCart.Payment.Total,
			&rawCart.Payment.Create_date, &rawCart.Payment.Modify_date, &rawMainCart.Id, &rawMainCart.Order_id, &rawMainCart.User_id, &rawMainCart.Total_amount,
			&rawMainCart.Payment_type, &rawMainCart.Create_date, &rawMainCart.Modify_date); err != nil {
			dataCart = append(dataCart, rawCart)
		}
		//totalAmount, _ := strconv.ParseFloat(rawCart.Payment.Total, 64)
		// total := rows.Scan(total)
		//total += totalAmount
		dataCart = append(dataCart, rawCart)
	}

	for rowsSum.Next() {
		total := rowsSum.Scan(&total)
		fmt.Println("Sum total", total)
	}
	if err != nil {
		log.Fatal(err)
	}
	// return e.JSON(http.StatusOK, dataCart)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"order": dataCart,
		"total": total,
	})

}
