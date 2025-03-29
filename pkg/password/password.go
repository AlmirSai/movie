package password

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// I think this bad realisation for production setup
	// TODO: Add salt in hash for password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), fmt.Errorf("error hashing password: %v", err)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
