package gostr

import "testing"

func TestSha256String(t *testing.T) {
	dataItems := []struct {
		Data   string
		Result string
	}{
		{"bismillah", "a1dd395c19f3a1f7ff0e6bfe3ee6028e4373b41085a6da776ab02a8baf129abb"},
	}
	for _, item := range dataItems {
		if hash := Sha256String(item.Data); hash != item.Result {
			t.Errorf("Beklenen: %v, Gelen: %v", item.Result, hash)
		}
	}
}

func TestMd5String(t *testing.T) {
	dataItems := []struct {
		Data   string
		Result string
	}{
		{"bismillah", "e172dd95f4feb21412a692e73929961e"},
	}
	for _, item := range dataItems {
		if hash := Md5String(item.Data); hash != item.Result {
			t.Errorf("Beklenen: %v, Gelen: %v", item.Result, hash)
		}
	}
}
