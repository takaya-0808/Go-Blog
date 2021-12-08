package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CreateHashPass(pass string) string {

	passhash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "0"
	}
	passHash := string(passhash)
	return passHash
}

func CheckHashPass(hashpass string, pass string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
