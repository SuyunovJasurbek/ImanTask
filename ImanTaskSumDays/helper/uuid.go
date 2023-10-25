package helper

import "github.com/google/uuid"

// generate uuid string
func GenerateUUID() string {
	return uuid.NewString()
}