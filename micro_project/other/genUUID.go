package other

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	// Generate a new UUID
	id := uuid.New()

	// Convert the UUID to a string representation
	uuidString := id.String()

	return uuidString
}
