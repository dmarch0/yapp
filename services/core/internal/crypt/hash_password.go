package crypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println("Error hasing password", err)
	}

	return string(hash)
}
