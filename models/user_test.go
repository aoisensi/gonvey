package models

import (
	"testing"
)

func TestUsernameCheck(t *testing.T) {
	sample := map[string]bool{
		"aoisensi":          true,
		"AOiSenSi":          true,
		"a":                 true,
		"_":                 true,
		"aaAAaAA___3443":    true,
		"1111222233334444":  true,
		"1111222233334444_": false,
		"":                  false,
		"テスト":               false,
		"@aoisensi":         false,
	}
	for name, ok := range sample {
		if CheckUsername(name) != ok {
			t.Fatalf("Failed \"%s\"", name)
		}
	}
}
