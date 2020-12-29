package gostr

import (
	"strings"
	"unicode"
)

func TurkceBuyukHarfeCevir(metin string) string {
	return strings.ToUpperSpecial(unicode.TurkishCase, metin)
}

func TurkceKucukHarfeCevir(metin string) string {
	return strings.ToLowerSpecial(unicode.TurkishCase, metin)
}
