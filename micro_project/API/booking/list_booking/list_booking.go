package listbooking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	bookstadium "project/models/book_stadium"
	"project/schemas"
	"time"

	"github.com/labstack/echo/v4"
)

func ListBookings(e echo.Context) error {
	body := &bookstadium.Booking{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	// Check if any parameter in the body is null or empty
	if body.User_id == "" {
		return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to marshal JSON data"})
	}
	fmt.Printf("Params: %s\n", jsonData)

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Query Booking for set existingBookings
	queryBook := "SELECT bk.*,k.brand_title,s.stadium_number,(s.price * bk.reservation_hours) AS totalPrice,COALESCE(sp.status, 0) AS status_payment  FROM booking bk inner join brand k on k.id = bk.brand_id inner join stadium s on s.id = bk.stadium_id LEFT JOIN stadium_payment sp ON sp.booking_id = bk.id WHERE bk.user_id=? order by CASE WHEN bk.del = 1 THEN 2 ELSE 1 END, bk.reservation_date desc, bk.start_time asc"
	rowsBook, err := db.Query(queryBook, body.User_id)
	if err != nil {
		fmt.Println("Error querying bookings:", err)
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to query existing bookings"})
	}
	defer rowsBook.Close()

	// Loop through the query results and save them to existingBookings
	bookings := []bookstadium.Booking{}
	for rowsBook.Next() {
		var booking bookstadium.Booking
		err := rowsBook.Scan(&booking.Id, &booking.User_id, &booking.Brand_id, &booking.Stadium_id, &booking.Reservation_hours, &booking.Reservation_date, &booking.Start_time, &booking.End_time, &booking.Create_date, &booking.Modify_date, &booking.Del, &booking.Status_payment, &booking.Brand_title, &booking.Stadium_number, &booking.Total_Price, &booking.Payment_success)
		if err != nil {
			fmt.Println("Error scanning booking:", err)
			return e.NoContent(http.StatusNoContent)

		}

		// Parse booking.Reservation_date and booking.Start_time to time.Time
		reservationDate, err := time.Parse("2006-01-02", booking.Reservation_date)
		if err != nil {
			return e.NoContent(http.StatusNoContent)
		}
		// fmt.Println("reservationDate => :", reservationDate)

		startTime, err := time.Parse("15:04", booking.Start_time)
		if err != nil {
			return e.NoContent(http.StatusNoContent)
		}

		// fmt.Println("startTime => :", startTime)
		// Check if the reservation is overdue
		currentTime := time.Now()
		currentTimeString := currentTime.Format("15")
		currentStartTimeString := startTime.Format("15")
		cutCurrentTime := currentTime.Truncate(24 * time.Hour)

		if reservationDate.After(cutCurrentTime) || (reservationDate.Equal(cutCurrentTime) && currentStartTimeString > currentTimeString) {
			booking.OverDue = "0"
		} else {
			booking.OverDue = "1"
		}
		bookings = append(bookings, booking)
	}
	fmt.Println("Number of bookings found:", len(bookings))
	return e.JSON(http.StatusOK, bookings)
}
