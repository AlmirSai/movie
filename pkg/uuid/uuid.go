package uuid

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.New().String()
}

func IsValidUUID(this string) bool {
	_, err := uuid.Parse(this)
	return err == nil
}
