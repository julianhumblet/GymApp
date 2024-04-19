package hashing

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %s", err)
	}

	hashedPassword := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashedPassword, nil
}
