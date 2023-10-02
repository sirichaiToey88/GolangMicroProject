package main

import (
	"project/API/booking"
	"project/API/booking/brand"
	listbooking "project/API/booking/list_booking"
	managebooking "project/API/booking/manage_booking"
	"project/API/customer"
	"project/API/login"
	"project/API/payment"
	"project/API/readQr"
	"project/API/shoppingCart"

	_ "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func authorizationMiddleware(c *gin.Context) {
// 	s := c.Request.Header.Get("Authorization")
// 	token := string.TrimPrefix(s, "Bearer")

// 	if err := validateToken(token); err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}
// }

func main() {
	// qrCodeValue := "0016A000000677010112011509940001649121802150034011941000040315201911260000095" // ค่า QR code ที่ต้องการแปลงเป็น JSON
	// result, err := other.ToJSON(qrCodeValue)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("JSON Result:", result)
	// e := echo.New()
	// e.Use(middleware.Logger())

	// e.GET("/getAll", customer.GetAllCustomerEcho)
	// e.POST("/register", customer.RegisterEcho)
	// e.POST("/updateProfile", customer.UpdateCustInfo)
	// e.POST("/getUser", customer.GetUserHandleByID)
	// e.POST("/login", login.LoginHandle)
	// e.Start(":8080")
	e := echo.New()
	e.Use(middleware.Logger())

	// Define the JWT middleware with a secret key
	// jwtMiddleware := middleware.JWT([]byte("Sirichai"))

	// Use the JWT middleware to protect the routes
	// e.GET("/getAll", customer.GetAllCustomerEcho, jwtMiddleware)
	// e.POST("/register", customer.RegisterEcho, jwtMiddleware)
	// e.POST("/updateProfile", customer.UpdateCustInfo, jwtMiddleware)
	// e.POST("/getUser", customer.GetUserHandleByID, jwtMiddleware)
	// e.POST("/shopping", shoppingCart.SaveShoppingCart, jwtMiddleware)
	// e.GET("/compareQR", readQr.CompareQR, jwtMiddleware)
	// e.POST("/SearchOrdeyByID", shoppingCart.SearchOrdeyByID, jwtMiddleware)
	// e.POST("/Findslot", booking.FindSlot, jwtMiddleware)
	// e.POST("/addNewBrand", brand.AddNewBrand, jwtMiddleware)
	// e.POST("/addStaduimToBrand", brand.AddStaduimToBrand)
	// e.GET("/listAllBrand", booking.ListAllBrand, jwtMiddleware)
	// e.POST("/AddBooking", managebooking.AddBooking, jwtMiddleware)
	// e.POST("/ListBooking", listbooking.ListBookings, jwtMiddleware)
	e.GET("/getAll", customer.GetAllCustomerEcho)
	e.POST("/register", customer.RegisterEcho)
	e.POST("/updateProfile", customer.UpdateCustInfo)
	e.POST("/getUser", customer.GetUserHandleByID)
	e.POST("/shopping", shoppingCart.SaveShoppingCart)
	e.GET("/compareQR", readQr.CompareQR)
	e.POST("/SearchOrdeyByID", shoppingCart.SearchOrdeyByID)
	e.POST("/Findslot", booking.FindSlot)
	e.POST("/addNewBrand", brand.AddNewBrand)
	e.POST("/addStaduimToBrand", brand.AddStaduimToBrand)
	e.POST("/AddNewBrandMobile", brand.AddNewBrandMobile)
	e.GET("/listAllBrand", booking.ListAllBrand)
	e.POST("/AddBooking", managebooking.AddBooking)
	e.POST("/PaymentBooking", managebooking.PaymentBooking)
	e.POST("/CancelBookings", managebooking.CancelBookings)
	e.POST("/ListBooking", listbooking.ListBookings)
	e.POST("/PaymentStripe", payment.AuthenAndPayment)
	e.POST("/PaymentOmise", payment.PaymentOmise)
	// e.POST("/shopping", shoppingCart.SaveShoppingCart)
	e.POST("/login", login.LoginHandle)

	e.Start(":8080")
}
