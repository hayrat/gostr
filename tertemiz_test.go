package gostr

import "testing"

func TestTertemiz(t *testing.T) {
	dataItems := []struct {
		Data   string
		Result string
	}{
		{"DİŞÇİ", "disci"},
		{"dişçi", "disci"},
		{"AĞAÇ", "agac"},
		{"ağaç", "agac"},
		{"ÖĞÜN", "ogun"},
		{"öğün", "ogun"},
		{"ILIK", "ilik"},
		{"ılık", "ilik"},
		{"sonda boşluk olmamalı ", "sonda-bosluk-olmamali"},
		{"emojiye hayir 😀😃😂😎", "emojiye-hayir"},
		{"Dişçi ağacın altında ılık süt içti", "disci-agacin-altinda-ilik-sut-icti"},
	}
	for _, item := range dataItems {
		if temiz := Tertemiz(item.Data); temiz != item.Result {
			t.Errorf("Beklenen: %v, Gelen: %v", item.Result, temiz)
		}
	}
}