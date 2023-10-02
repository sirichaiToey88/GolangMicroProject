package booking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	bookstadium "project/models/book_stadium"
	"project/schemas"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
)

// ReservationSlot represents a time slot for table reservation
type ReservationSlot struct {
	Id                string `json:"id"`
	User_id           string `json:"user_id"`
	Brand_id          string `json:"brand_id"`
	Stadium_id        string `json:"stadium_id"`
	Reservation_hours string `json:"reservation_hours"`
	Reservation_date  string `json:"reservation_date"`
	Date              string `json:"date"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	Create_date       string `json:"create_date"`
	Modify_date       string `json:"modify_date"`
	Del               string `json:"del"`
	Status_payment    string `json:"status_payment"`
}

// FormattedReservationSlot represents a formatted time slot for JSON response
type FormattedReservationSlot struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// RequestParams represents the request parameters for finding available slots
type RequestParams struct {
	ReservationDate  string `json:"reservationDate"`
	ReservationHours int    `json:"reservationHours"`
	Brand_id         string `json:"brand_id"`
	Stadium_id       string `json:"stadium_id"`
}

func FindSlot(c echo.Context) error {
	// Parse the request parameters
	var params RequestParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request parameters"})
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to marshal JSON data"})
	}
	fmt.Printf("Params: %s\n", jsonData)

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Query Brand for get time_open and time_close
	rBrand := bookstadium.Brand{}
	query := "SELECT * FROM brand WHERE id=?"

	row := db.QueryRow(query, params.Brand_id)

	err = row.Scan(&rBrand.Id, &rBrand.Brand_title, &rBrand.Image_url, &rBrand.Time_open, &rBrand.Time_close, &rBrand.Location, &rBrand.Tell, &rBrand.Address, &rBrand.Create_date, &rBrand.Modify_date)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Brand ID not found.")
	}

	// Dummy existing bookings for demonstration purposes
	existingBookings := []ReservationSlot{}

	//Query Booking for set existingBookings
	queryBook := "SELECT * FROM booking WHERE brand_id=? and stadium_id=? and reservation_date=? and del = 0"

	rowsBook, err := db.Query(queryBook, params.Brand_id, params.Stadium_id, params.ReservationDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to query existing bookings"})
	}
	defer rowsBook.Close()

	// Loop through the query results and save them to existingBookings
	for rowsBook.Next() {
		var booking ReservationSlot
		err := rowsBook.Scan(&booking.Id, &booking.User_id, &booking.Brand_id, &booking.Stadium_id, &booking.Reservation_hours, &booking.Reservation_date, &booking.StartTime, &booking.EndTime, &booking.Create_date, &booking.Modify_date, &booking.Del, &booking.Status_payment)
		if err != nil {
			fmt.Println("Error scanning booking:", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to scan existing bookings"})
		}
		existingBookings = append(existingBookings, booking)
		fmt.Println("Booking ID:", booking.Id, "User ID:", booking.User_id, "Start Time:", booking.StartTime, "End Time:", booking.EndTime)
	}

	// End Query Booking for set existingBookings
	// Set the opening and closing time for every day
	layoutTime := "15:04"
	openingTime, _ := time.Parse(layoutTime, rBrand.Time_open)
	closingTime, _ := time.Parse(layoutTime, rBrand.Time_close)

	// Parse the reservation date string into a time.Time object
	layout := "2006-01-02"
	reservationDate, err := time.Parse(layout, params.ReservationDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid reservation date format"})
	}

	// Get the current time
	now := time.Now()
	if (reservationDate.Before(now) && reservationDate.Day() != now.Day()) || (now.Before(closingTime) && reservationDate.Day() == now.Day()) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Not booking before current date"})
	}

	// Calculate the reservation start time and end time
	reservationStartTime := openingTime
	reservationDuration := time.Duration(params.ReservationHours) * time.Hour
	reservationEndTime := reservationStartTime.Add(reservationDuration)
	// Calculate the available reservation time slots

	if len(existingBookings) != 0 && params.ReservationHours == 3 && reservationDate.Day() != now.Day() {
		findAvailableSlots := findAvailableSlots(reservationStartTime, reservationEndTime, closingTime, params.ReservationDate, params.ReservationHours, existingBookings)
		// Convert the ReservationSlot objects to FormattedReservationSlot for JSON response
		formattedSlots := make([]FormattedReservationSlot, len(findAvailableSlots))
		for i, slot := range findAvailableSlots {
			formattedSlots[i] = FormattedReservationSlot{
				StartTime: slot.StartTime,
				EndTime:   slot.EndTime,
			}
		}
		// Return the JSON response with the available time slots
		return c.JSON(http.StatusOK, formattedSlots)
	} else {
		availableSlots := getAvailableReservationSlots(existingBookings, openingTime, closingTime, reservationStartTime, reservationEndTime, params.ReservationDate, params.ReservationHours)
		// Convert the ReservationSlot objects to FormattedReservationSlot for JSON response
		formattedSlots := make([]FormattedReservationSlot, len(availableSlots))
		for i, slot := range availableSlots {
			formattedSlots[i] = FormattedReservationSlot{
				StartTime: slot.StartTime,
				EndTime:   slot.EndTime,
			}
		}
		// Return the JSON response with the available time slots
		if len(formattedSlots) != 0 {
			return c.JSON(http.StatusOK, formattedSlots)
		} else {
			emptySlots := []FormattedReservationSlot{}
			return c.JSON(http.StatusNoContent, echo.Map{"slots": emptySlots})
		}

	}
}

// getAvailableReservationSlots calculates the available reservation time slots
func getAvailableReservationSlots(existingBookings []ReservationSlot, openingTime, closingTime, reservationStartTime, reservationEndTime time.Time, reservationDate string, reservationDuration int) []ReservationSlot {
	var availableSlots []ReservationSlot
	// Check if the reservation start time is before the opening time
	if reservationStartTime.Before(openingTime) {
		reservationStartTime = openingTime
	}

	// Check if the reservation end time is after the closing time
	if reservationEndTime.After(closingTime) {
		reservationEndTime = closingTime
	}

	getNowTime := time.Now()

	// Convert the current time to your desired format "15:04"
	endSlot := getNowTime.Format("15")
	dateNow := time.Now().Format("2006-01-02")
	// layoutTime := "15:04"

	// Sort the reservations list based on the start_time field
	sort.Slice(existingBookings, func(i, j int) bool {
		return existingBookings[i].StartTime < existingBookings[j].StartTime
	})

	if endSlot == reservationEndTime.Add(time.Hour).Format("15") && dateNow >= reservationDate {
		reservationCurrEndTime := reservationEndTime.Add(time.Hour)

		reservationStartTime = reservationCurrEndTime
		reservationDuration := time.Duration(reservationDuration) * time.Hour
		reservationEndTime = reservationStartTime.Add(reservationDuration)
	}
	// Parse the date string into a time.Time object
	layoutDate := "2006-01-02"
	date, _ := time.Parse(layoutDate, reservationDate)

	for startTime := reservationStartTime; startTime.Add(reservationEndTime.Sub(reservationStartTime)).Before(closingTime); startTime = startTime.Add(reservationEndTime.Sub(reservationStartTime)) {
		available := true
		layoutTime := "15:04"
		dateNow := time.Now().Format("2006-01-02")
		checkStartTimeWantBookDate := time.Now().Format("2006-01-02") + " " + startTime.Format("15")
		checkStartTimeCurrentDate := time.Now().Format("2006-01-02 15")

		// Compare the date and time separately for each existing booking
		for _, existingBooking := range existingBookings {
			existingStartTime, _ := time.Parse(layoutTime, existingBooking.StartTime)
			existingEndTime, _ := time.Parse(layoutTime, existingBooking.EndTime)
			existingDate, _ := time.Parse(layoutDate, existingBooking.Reservation_date)

			reservedStartTime := startTime
			reservedEndTime := startTime.Add(reservationEndTime.Sub(reservationStartTime))

			if dateNow >= reservationDate {
				if checkStartTimeWantBookDate < checkStartTimeCurrentDate {
					// There is an overlap with an existing booking
					available = false
					break
				}
			}

			if existingDate.Equal(date) && existingStartTime.Before(reservedEndTime) && existingEndTime.After(reservedStartTime) {
				// There is an overlap with an existing booking
				available = false
				break
			}

		}

		if dateNow >= reservationDate {
			if checkStartTimeWantBookDate > checkStartTimeCurrentDate {
				if available {
					availableSlots = append(availableSlots, ReservationSlot{
						Date:      startTime.Format(layoutDate),
						StartTime: startTime.Format(layoutTime),
						EndTime:   startTime.Add(reservationEndTime.Sub(reservationStartTime)).Format(layoutTime),
					})
				}
			}
		} else {
			if available {
				availableSlots = append(availableSlots, ReservationSlot{
					Date:      startTime.Format(layoutDate),
					StartTime: startTime.Format(layoutTime),
					EndTime:   startTime.Add(reservationEndTime.Sub(reservationStartTime)).Format(layoutTime),
				})
			}
		}
	}

	return availableSlots
}

// func getAvailableReservationSlots(existingBookings []ReservationSlot, openingTime, closingTime, reservationStartTime, reservationEndTime time.Time, reservationDate string, reservationDuration int) []ReservationSlot {
func findAvailableSlots(reservationStartTime, reservationEndTime, closingTime time.Time, reservationDate string, reservationDuration int, existingBookings []ReservationSlot) []ReservationSlot {
	availableSlots := []ReservationSlot{}
	layoutTime := "15:04"
	layoutDate := "2006-01-02"
	reservationDurationInHours := time.Duration(reservationDuration) * time.Hour

	// Sort existingBookings by StartTime
	sort.Slice(existingBookings, func(i, j int) bool {
		return existingBookings[i].StartTime < existingBookings[j].StartTime
	})

	for i := 0; i < len(existingBookings)-1; i++ {
		endTime1, _ := time.Parse(layoutTime, existingBookings[i].EndTime)
		startTime2, _ := time.Parse(layoutTime, existingBookings[i+1].StartTime)

		gap := startTime2.Sub(endTime1)
		if gap >= reservationDurationInHours {
			// There is a gap between the end time of the current slot and the start time of the next slot
			// Create a new slot starting from the end time of the current slot
			startTime := endTime1
			endTime := endTime1.Add(reservationDurationInHours)

			// Check if the new slot overlaps with any existing bookings
			available := true
			for _, existingBooking := range existingBookings {
				existingStartTime, _ := time.Parse(layoutTime, existingBooking.StartTime)
				existingEndTime, _ := time.Parse(layoutTime, existingBooking.EndTime)
				existingDate, _ := time.Parse(layoutDate, existingBooking.Reservation_date)
				reservationDate, _ := time.Parse(layoutDate, existingBooking.Reservation_date)

				if reservationDate.Equal(existingDate) && ((startTime.Before(existingEndTime) && endTime.After(existingStartTime)) ||
					(startTime.Equal(existingStartTime) || endTime.Equal(existingEndTime))) {
					available = false
					break
				}
				if startTime.Before(time.Now()) {
					// หาก startTime น้อยกว่าเวลาปัจจุบัน แสดงว่า Slot นี้ไม่สามารถจองได้
					available = false
					break
				}
			}

			if available {
				availableSlots = append(availableSlots, ReservationSlot{
					Date:      startTime.Format(layoutDate),
					StartTime: startTime.Format(layoutTime),
					EndTime:   endTime.Format(layoutTime),
				})
			}
		}
	}

	// Check if there is a gap between the last end time and the closing time
	lastEndTime, _ := time.Parse(layoutTime, existingBookings[len(existingBookings)-1].EndTime)
	gap := closingTime.Sub(lastEndTime)
	if gap >= reservationDurationInHours {
		// Create a new slot starting from the last end time
		startTime := lastEndTime
		endTime := lastEndTime.Add(reservationDurationInHours)

		// Check if the new slot overlaps with any existing bookings
		available := true
		for _, existingBooking := range existingBookings {
			existingStartTime, _ := time.Parse(layoutTime, existingBooking.StartTime)
			existingEndTime, _ := time.Parse(layoutTime, existingBooking.EndTime)
			reservationDate, _ := time.Parse(layoutDate, existingBooking.Reservation_date)
			existingDate, _ := time.Parse(layoutDate, existingBooking.Reservation_date)

			if reservationDate.Equal(existingDate) && ((startTime.Before(existingEndTime) && endTime.After(existingStartTime)) ||
				(startTime.Equal(existingStartTime) || endTime.Equal(existingEndTime))) {
				available = false
				break
			}
		}

		if available {
			availableSlots = append(availableSlots, ReservationSlot{
				Date:      startTime.Format(layoutDate),
				StartTime: startTime.Format(layoutTime),
				EndTime:   endTime.Format(layoutTime),
			})
		}

		// Continue creating slots until the closing time is reached
		for endTime.Add(reservationDurationInHours).Before(closingTime) {
			startTime = endTime
			endTime = endTime.Add(reservationDurationInHours)

			// Check if the new slot overlaps with any existing bookings
			available := true
			for _, existingBooking := range existingBookings {
				existingStartTime, _ := time.Parse(layoutTime, existingBooking.StartTime)
				existingEndTime, _ := time.Parse(layoutTime, existingBooking.EndTime)
				reservationDate, _ := time.Parse(layoutDate, existingBooking.Reservation_date)
				existingDate, _ := time.Parse(layoutDate, existingBooking.Reservation_date)

				if reservationDate.Equal(existingDate) && ((startTime.Before(existingEndTime) && endTime.After(existingStartTime)) ||
					(startTime.Equal(existingStartTime) || endTime.Equal(existingEndTime))) {
					available = false
					break
				}
			}

			if available {
				availableSlots = append(availableSlots, ReservationSlot{
					Date:      startTime.Format(layoutDate),
					StartTime: startTime.Format(layoutTime),
					EndTime:   endTime.Format(layoutTime),
				})
			}
		}
	}

	return availableSlots
}
