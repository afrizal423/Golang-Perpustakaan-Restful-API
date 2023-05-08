package utils

import (
	"log"

	"github.com/alexedwards/argon2id"
)

type Hash struct {
}

func (h *Hash) HashPassword(pw string) string {
	hash, err := argon2id.CreateHash(pw, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func (h *Hash) VerifikasiPassword(pw string, hashedPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(pw, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return match, err
}
