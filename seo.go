package gostr

import (
	"strings"
	"unicode"
)

func Seo(metin string) string {
	metin = TurkceKucukHarfeCevir(metin)
	metin = Latinize(metin)
	metin = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '-'
	}, metin)
	return strings.Trim(metin, "-")
}
