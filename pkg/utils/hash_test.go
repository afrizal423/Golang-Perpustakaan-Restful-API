package utils

import (
	"regexp"
	"strings"
	"testing"
)

func TestHashPasswordh(t *testing.T) {
	hashRX, err := regexp.Compile(`^\$argon2id\$v=19\$m=65536,t=1,p=2\$[A-Za-z0-9+/]{22}\$[A-Za-z0-9+/]{43}$`)
	if err != nil {
		t.Fatal(err)
	}
	var hash Hash
	hash1 := hash.HashPassword("pa$$word")

	if !hashRX.MatchString(hash1) {
		t.Errorf("hash %q not in correct format", hash1)
	}

	hash2 := hash.HashPassword("pa$$word")

	if strings.Compare(hash1, hash2) == 0 {
		t.Error("hashes must be unique")
	}
}

func TestVerifikasiPassword(t *testing.T) {
	var hash Hash
	pw := hash.HashPassword("pa$$word")

	match, err := hash.VerifikasiPassword("pa$$word", pw)
	if err != nil {
		t.Fatal(err)
	}

	if !match {
		t.Error("expected password and hash to match")
	}

	match, err = hash.VerifikasiPassword("otherPa$$word", pw)
	if err != nil {
		t.Fatal(err)
	}

	if match {
		t.Error("expected password and hash to not match")
	}
}
