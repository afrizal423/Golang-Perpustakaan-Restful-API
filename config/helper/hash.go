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
