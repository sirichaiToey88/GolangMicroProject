package managebooking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	bookstadium "project/models/book_stadium"
	"project/other"
	"project/schemas"

	"github.com/labstack/echo/v4"
)

func AddBooking(e echo.Context) error {
	// Parse the request body to extract the data to be inserted
	body := []bookstadium.Booking{}
	// if err := e.Bind(body); err != nil {
	// 	return e.JSON(http.StatusBadRequest, err)
	// }

	if err := e.Bind(&body); err != nil {
		jsonData, _ := json.Marshal(body)
		log.Println("Failed to marshal JSON: %s", jsonData)
		log.Println(string(jsonData))
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
	}

	requestJSON, _ := json.Marshal(e.Request().Body)
	fmt.Println("Received Request:", string(requestJSON))

	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return nil
	}

	// Print the JSON-formatted string
	fmt.Println(string(jsonData))

	// Check if any parameter in the body is null or empty
	for _, booking := range body {
		if booking.End_time == "" || booking.Start_time == "" || booking.Stadium_id == "" || booking.Brand_id == "" || booking.User_id == "" || booking.Reservation_date == "" || booking.Reservation_hours == "" {
			return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
		}

		db, err := schemas.Connectdb()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		//Query Brand for get time_open and time_close
		rBrand := bookstadium.Brand{}
		query := "SELECT * FROM brand WHERE id=?"

		row := db.QueryRow(query, booking.Brand_id)

		errBrand := row.Scan(&rBrand.Id, &rBrand.Brand_title, &rBrand.Image_url, &rBrand.Time_open, &rBrand.Time_close, &rBrand.Location, &rBrand.Tell, &rBrand.Address, &rBrand.Create_date, &rBrand.Modify_date)

		if errBrand != nil {
			return e.JSON(http.StatusNotFound, "Brand ID not found.")
		}

		//Query Stadium for get status
		rStadiumNumber := bookstadium.Stadium{}
		queryStadiumNumber := "SELECT * FROM stadium WHERE id=?"

		rowStadiumNumber := db.QueryRow(queryStadiumNumber, booking.Stadium_id)
		err = rowStadiumNumber.Scan(&rStadiumNumber.Id, &rStadiumNumber.Brand_id, &rStadiumNumber.Type_stadium, &rStadiumNumber.Status, &rStadiumNumber.Image_url, &rStadiumNumber.Promotion, &rStadiumNumber.Price, &rStadiumNumber.Create_date, &rStadiumNumber.Modify_date, &rStadiumNumber.Stadium_number)

		if err != nil {
			return e.JSON(http.StatusInternalServerError, err.Error())
		}

		if errBrand != nil {
			return e.JSON(http.StatusNotFound, "Brand ID not found.")
		} else if rStadiumNumber.Status != "0" {
			return e.JSON(http.StatusInternalServerError, "Stadium is busy.")
		} else {
			stmt, err := db.Prepare("INSERT INTO booking(id,User_id,Brand_id,Stadium_id,Reservation_hours,Reservation_date,Start_time,End_time,Create_date,modify_date,del,status_payment) VALUES(?,?,?,?,?,?,?,?,?,?,?,?);")

			if err != nil {
				return e.JSON(http.StatusInternalServerError, err.Error())
			}
			defer stmt.Close()

			//Update status stadium
			// upStatus, err := db.Prepare((`UPDATE Stadium SET status = "1" where id=?`))

			// if err != nil {
			// 	return e.JSON(http.StatusInternalServerError, err.Error())
			// }
			// defer upStatus.Close()

			// _, err = upStatus.Exec(body.Stadium_id)

			// if err != nil {
			// 	return e.JSON(http.StatusInternalServerError, err.Error())
			// }

			// Execute the INSERT statement
			// UUID := other.GenerateUUID()
			UUIDBooking := other.GenerateUUID()
			Createdate := other.DateTime()
			Modifydate := other.DateTime()

			result, err := stmt.Exec(UUIDBooking, booking.User_id, booking.Brand_id, booking.Stadium_id, booking.Reservation_hours, booking.Reservation_date, booking.Start_time, booking.End_time, Createdate, Modifydate, "0", "0")

			// if booking.NeedToPay == "1" {
			// 	// Insert Payment
			// 	stmtPay, err := db.Prepare("INSERT INTO stadium_payment(id,User_id,Booking_id,Status,Image_url,Total,Create_date,modify_date,source,destination) VALUES(?,?,?,?,?,?,?,?,?,?);")
			// 	_, err = stmtPay.Exec(UUID, booking.User_id, booking.Id, booking.Payment.Status, booking.Payment.Image_url, booking.Payment.Total, Createdate, Modifydate, booking.Payment.Source, booking.Payment.Destination)

			// 	if err != nil {
			// 		return e.JSON(http.StatusInternalServerError, err.Error())
			// 	}
			// 	defer stmtPay.Close()
			// }

			if err != nil {
				return e.JSON(http.StatusInternalServerError, err.Error())
			}
			defer stmt.Close()

			if err != nil {
				return e.JSON(http.StatusInternalServerError, err.Error())
			} else {
				id, _ := result.LastInsertId()
				booking.Id = string(id)
				return e.JSON(http.StatusOK, map[string]interface{}{
					"Status":     200,
					"BookDetail": body,
				})
			}
		}
	}
	return e.JSON(http.StatusOK, "200")
}
