package booking

// package booking

// import (
// 	"net/http"
// 	"time"

// 	"github.com/labstack/echo/v4"
// )

// // ReservationSlot represents a time slot for table reservation
// type ReservationSlot struct {
// 	Date      string `json:"date"`
// 	StartTime string `json:"start_time"`
// 	EndTime   string `json:"end_time"`
// }

// // FormattedReservationSlot represents a formatted time slot for JSON response
// type FormattedReservationSlot struct {
// 	StartTime string `json:"start_time"`
// 	EndTime   string `json:"end_time"`
// }

// // RequestParams represents the request parameters for finding available slots
// type RequestParams struct {
// 	OpeningTime      string `json:"opening_time"`
// 	ClosingTime      string `json:"closing_time"`
// 	ReservationDate  string `json:"reservationDate"`
// 	ReservationHours int    `json:"reservationHours"`
// }

// func FindSlot(c echo.Context) error {
// 	// Parse the request parameters
// 	var params RequestParams
// 	if err := c.Bind(&params); err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request parameters"})
// 	}

// 	// Set the opening and closing time for every day
// 	layoutTime := "15:04"
// 	openingTime, _ := time.Parse(layoutTime, params.OpeningTime)
// 	closingTime, _ := time.Parse(layoutTime, params.ClosingTime)

// 	// Parse the reservation date string into a time.Time object
// 	layout := "2006-01-02"
// 	reservationDate, err := time.Parse(layout, params.ReservationDate)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid reservation date format"})
// 	}

// 	// Get the current time
// 	now := time.Now()
// 	if (reservationDate.Before(now) && reservationDate.Day() != now.Day()) || (now.Before(closingTime) && reservationDate.Day() == now.Day()) {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Not booking before current date"})
// 	}

// 	// Calculate the reservation start time and end time
// 	reservationStartTime, _ := time.Parse(layoutTime, params.OpeningTime)
// 	reservationDuration := time.Duration(params.ReservationHours) * time.Hour
// 	reservationEndTime := reservationStartTime.Add(reservationDuration)

// 	// Dummy existing bookings for demonstration purposes
// 	existingBookings := []ReservationSlot{
// 		{Date: "2023-07-23", StartTime: "12:00", EndTime: "14:00"},
// 		{Date: "2023-07-25", StartTime: "14:00", EndTime: "16:00"},
// 		{Date: "2023-07-26", StartTime: "12:00", EndTime: "14:00"},
// 	}

// 	// Calculate the available reservation time slots
// 	availableSlots := getAvailableReservationSlots(existingBookings, openingTime, closingTime, reservationStartTime, reservationEndTime, params.ReservationDate)

// 	// Convert the ReservationSlot objects to FormattedReservationSlot for JSON response
// 	formattedSlots := make([]FormattedReservationSlot, len(availableSlots))
// 	for i, slot := range availableSlots {
// 		formattedSlots[i] = FormattedReservationSlot{
// 			StartTime: slot.StartTime,
// 			EndTime:   slot.EndTime,
// 		}
// 	}

// 	// Return the JSON response with the available time slots
// 	return c.JSON(http.StatusOK, formattedSlots)
// }

// // getAvailableReservationSlots calculates the available reservation time slots
// func getAvailableReservationSlots(existingBookings []ReservationSlot, openingTime, closingTime, reservationStartTime, reservationEndTime time.Time, reservationDate string) []ReservationSlot {
// 	var availableSlots []ReservationSlot

// 	// Check if the reservation start time is before the opening time
// 	if reservationStartTime.Before(openingTime) {
// 		reservationStartTime = openingTime
// 	}

// 	// Check if the reservation end time is after the closing time
// 	if reservationEndTime.After(closingTime) {
// 		reservationEndTime = closingTime
// 	}

// 	// Parse the date string into a time.Time object
// 	layoutDate := "2006-01-02"
// 	date, _ := time.Parse(layoutDate, reservationDate)

// 	// Check for overlapping with existing bookings
// 	for startTime := reservationStartTime; startTime.Add(reservationEndTime.Sub(reservationStartTime)).Before(closingTime); startTime = startTime.Add(reservationEndTime.Sub(reservationStartTime)) {
// 		available := true

// 		// Convert the start and end times of existing bookings to time.Time objects
// 		layoutTime := "15:04"

// 		// Compare the date and time separately for each existing booking
// 		for _, existingBooking := range existingBookings {
// 			existingStartTime, _ := time.Parse(layoutTime, existingBooking.StartTime)
// 			existingEndTime, _ := time.Parse(layoutTime, existingBooking.EndTime)
// 			existingDate, _ := time.Parse(layoutDate, existingBooking.Date)

// 			reservedStartTime := startTime
// 			reservedEndTime := startTime.Add(reservationEndTime.Sub(reservationStartTime))

// 			if existingDate.Equal(date) && existingStartTime.Before(reservedEndTime) && existingEndTime.After(reservedStartTime) {
// 				// There is an overlap with an existing booking
// 				available = false
// 				break
// 			}
// 		}

// 		// If the time slot is available, add it to the availableSlots list
// 		if available {
// 			availableSlots = append(availableSlots, ReservationSlot{
// 				Date:      startTime.Format(layoutDate),
// 				StartTime: startTime.Format(layoutTime),
// 				EndTime:   startTime.Add(reservationEndTime.Sub(reservationStartTime)).Format(layoutTime),
// 			})
// 		}
// 	}

// 	return availableSlots
// }
