package repositories

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GenerateId() string {
	return uuid.New().String()
}

func HashPassword(plainPassword string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return string(hashed)
}
