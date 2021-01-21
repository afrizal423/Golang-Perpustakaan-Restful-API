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
	if strings.Contains(err, "jenis_buku") {
		return errors.New("Jenis buku sudah terpakai")
	}
	if strings.Contains(err, "penulis_buku") {
		return errors.New("Penulis buku sudah terpakai")
	}
	if strings.Contains(err, "penerbit_buku") {
		return errors.New("Penerbit buku sudah terpakai")
	}
	if strings.Contains(err, "isbn") {
		return errors.New("ISBN sudah terpakai")
	}
	return errors.New("Incorrect Details")
}
