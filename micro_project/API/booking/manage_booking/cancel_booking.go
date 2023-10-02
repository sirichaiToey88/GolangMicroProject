package managebooking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	bookstadium "project/models/book_stadium"
	"project/schemas"

	"github.com/labstack/echo/v4"
)

func CancelBookings(e echo.Context) error {
	body := &bookstadium.Booking{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	// Check if any parameter in the body is null or empty
	if body.User_id == "" || body.Id == "" {
		fmt.Printf("One or more parameters are empty.")
		return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Params cancel 001 : %s\n", jsonData)
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to marshal JSON data"})
	}
	fmt.Printf("Params cancel: %s\n", jsonData)

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Query Booking for set existingBookings
	queryBook := "SELECT * FROM booking WHERE user_id=? and id=?"
	rowsBook, err := db.Query(queryBook, body.User_id, body.Id)
	if err != nil {
		fmt.Println("Error querying bookings:", err)
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to query existing bookings"})
	}
	defer rowsBook.Close()

	// Loop through the query results and save them to existingBookings
	bookings := []bookstadium.Booking{}

	for rowsBook.Next() {
		var booking bookstadium.Booking
		err := rowsBook.Scan(&booking.Id, &booking.User_id, &booking.Brand_id, &booking.Stadium_id, &booking.Reservation_hours, &booking.Reservation_date, &booking.Start_time, &booking.End_time, &booking.Create_date, &booking.Modify_date, &booking.Del, &booking.Status_payment)
		if err != nil {
			fmt.Println("Error scanning booking:", err)
			return e.NoContent(http.StatusNoContent)

		}

		bookings = append(bookings, booking)
	}

	// Check if a booking was found
	if len(bookings) == 0 {
		return e.JSON(http.StatusNotFound, echo.Map{"error": "booking not found"})
	}

	// Update the booking by setting del to 1
	updateQuery := "UPDATE booking SET del = 1 WHERE user_id=? AND id=?"
	_, err = db.Exec(updateQuery, body.User_id, body.Id)
	if err != nil {
		fmt.Println("Error updating booking:", err)
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to update booking"})
	}

	fmt.Println("Booking canceled:", body.Id)
	return e.JSON(http.StatusOK, echo.Map{"message": "Booking canceled successfully"})
}
