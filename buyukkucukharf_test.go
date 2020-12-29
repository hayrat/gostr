package gostr

import "testing"

func TestTurkceBuyukHarfeCevir(t *testing.T) {
	dataItems := []struct {
		Data   string
		Result string
	}{
		{"ılık", "ILIK"},
		{"dişçi", "DİŞÇİ"},
	}
	for _, item := range dataItems {
		if hash := TurkceBuyukHarfeCevir(item.Data); hash != item.Result {
			t.Errorf("Beklenen: %v, Gelen: %v", item.Result, hash)
		}
	}
}

func TestTurkceKucukHarfeCevir(t *testing.T) {
	dataItems := []struct {
		Data   string
		Result string
	}{
		{"ILık", "ılık"},
		{"DİŞÇİ", "dişçi"},
	}
	for _, item := range dataItems {
		if hash := TurkceKucukHarfeCevir(item.Data); hash != item.Result {
			t.Errorf("Beklenen: %v, Gelen: %v", item.Result, hash)
		}
	}
}
