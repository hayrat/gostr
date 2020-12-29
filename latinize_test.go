package gostr

import "testing"

func TestLatinize(t *testing.T) {
	dataItems := []struct {
		Data   string
		Result string
	}{
		{"DİŞÇİ", "DISCI"},
		{"dişçi", "disci"},
		{"AĞAÇ", "AGAC"},
		{"ağaç", "agac"},
		{"ÖĞÜN", "OGUN"},
		{"öğün", "ogun"},
		{"ILIK", "ILIK"},
		{"ılık", "ilik"},
	}
	for _, item := range dataItems {
		if hash := Latinize(item.Data); hash != item.Result {
			t.Errorf("Beklenen: %v, Gelen: %v", item.Result, hash)
		}
	}
}
