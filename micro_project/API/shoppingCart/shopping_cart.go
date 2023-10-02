package shoppingCart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	shoppingcart "project/models/shopping_cart"
	"project/other"
	"project/schemas"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveShoppingCart(e echo.Context) error {
	var body []shoppingcart.Shopping

	if err := e.Bind(&body); err != nil {
		jsonData, _ := json.Marshal(body)
		log.Println("Failed to marshal JSON: %s", jsonData)
		log.Println(string(jsonData))
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
	}

	//Log Data
	for _, order := range body {
		b, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return err
		}
		fmt.Println(string(b))
	}
	//End Log Data

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	UUID := other.GenerateUUID()
	Createdate := other.DateTime()
	Modifydate := other.DateTime()
	OrderID := other.RandomStringWithNumbers(10)

	insShopping, err := db.Prepare("INSERT INTO shopping_cart(id,user_id,product_id,product_title,Product_price,total_amount,quantity_product,payment_type,order_id,create_date,modify_date,image_url) VALUES(?,?,?,?,?,?,?,?,?,?,?,?);")

	total := 0.0
	for i := 0; i < len(body); i++ {
		UUId := other.GenerateUUID()
		totalAmount, err := strconv.ParseFloat(body[i].Payment.Total, 64)
		total += totalAmount
		// fmt.Println("Response total =>", total)

		if err != nil {
			return e.JSON(http.StatusInternalServerError, err.Error())
		}
		defer insShopping.Close()

		_, err = insShopping.Exec(UUId, body[i].User_id, body[i].Product_id, body[i].Product_title, body[i].Product_price, body[i].Total_amount, body[i].Quantity_product, body[i].Payment_type, OrderID, Createdate, Modifydate, body[i].Image_url)
	}

	insMainCart, err := db.Prepare("INSERT INTO main_cart(id,user_id,order_id,total_amount,payment_type,create_date,modify_date) VALUES(?,?,?,?,?,?,?);")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	defer insMainCart.Close()
	_, err = insMainCart.Exec(UUID, body[0].User_id, OrderID, total, body[0].Payment_type, Createdate, Modifydate)

	insPaymnet, err := db.Prepare("INSERT INTO payment(id,order_id,account,source,distination,total,create_date,modify_date) VALUES(?,?,?,?,?,?,?,?);")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	defer insPaymnet.Close()

	_, err = insPaymnet.Exec(UUID, OrderID, body[0].Payment.Account, body[0].Payment.Source, body[0].Payment.Distination, total, Createdate, Modifydate)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return e.JSON(http.StatusCreated, map[string]interface{}{
			"ShoppingCart": 200,
			"Payment":      200,
			"MainCart":     200,
			"Status":       200,
		})
	}
}
