package common

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(subject string) string {
	password := []byte(subject)

	passwordHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(passwordHash)
}

func ComparePasswordHash(subject string, object string) bool {
	passwordHash := []byte(subject)
	passwordStored := []byte(object)
	return bcrypt.CompareHashAndPassword(passwordHash, passwordStored) == nil
}
