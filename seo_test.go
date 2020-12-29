package gostr

import "testing"

func TestSeo(t *testing.T) {
	dataItems := []struct {
		Data   string
		Result string
	}{
		{"kofte123", "kofte123"},
		{"Siz Hayat Süren Köfteler Sizi kim pişirecek?",
			"siz-hayat-suren-kofteler-sizi-kim-pisirecek"},
	}
	for _, item := range dataItems {
		if hash := Seo(item.Data); hash != item.Result {
			t.Errorf("Beklenen: %v, Gelen: %v", item.Result, hash)
		}
	}
}
