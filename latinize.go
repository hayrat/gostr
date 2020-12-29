package gostr

import (
	"strings"
)

func Latinize(metin string) string {
	r := strings.NewReplacer(
		"Ğ", "G", "ğ", "g",
		"Ü", "U", "ü", "u",
		"İ", "I", "ı", "i",
		"Ş", "S", "ş", "s",
		"Ö", "O", "ö", "o",
		"Ç", "C", "ç", "c")

	return r.Replace(metin)
}
