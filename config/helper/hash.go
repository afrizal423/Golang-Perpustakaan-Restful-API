package helper

import (
	"log"

	"github.com/alexedwards/argon2id"
)

func Hash(pw string) string {
	password, err := argon2id.CreateHash(pw, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return password
}

func VerifikasiPassword(hashedPassword, password string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return match, err
}
