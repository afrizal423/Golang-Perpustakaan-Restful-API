package helper

import "strconv"

func StringkeInt(kata string) int {
	var angka int
	angka, _ = strconv.Atoi(kata)
	return angka
}
