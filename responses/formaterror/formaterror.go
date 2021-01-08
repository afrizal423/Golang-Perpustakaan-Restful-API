package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("Username sudah terpakai")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email sudah terpakai")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title sudah terpakai")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Password Salah")
	}
	return errors.New("Incorrect Details")
}
