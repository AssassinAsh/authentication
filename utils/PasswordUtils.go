package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//HashAndSaltPassword -
func HashAndSaltPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword(GetStringHash(pwd), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

//ComparePasswords -
func ComparePasswords(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword(GetStringHash(hashedPassword), GetStringHash(plainPassword))

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
