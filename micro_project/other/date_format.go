package other

import (
	"time"
)

func DateTime() string {
	currentTime := time.Now()

	// Format the datetime using a custom layout
	formattedDateTime := currentTime.Format("2006-01-02 15:04:05")

	return formattedDateTime
}
