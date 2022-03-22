package go_colorhash

import (
	"testing"
)

func TestHashBytes(t *testing.T) {
	testBytes := []struct {
		runAsUser bool
	}{}
	_ = testBytes
}
func TestHashString(t *testing.T) {
	testStrings := []struct {
		String string
		Value  int
		ID     string
	}{{String: "", Value: 7602086723416769149, ID: "Empty string"},
		{String: "123", Value: 1606385479620709231, ID: "123"},
		{String: "it's as easy as", Value: 5377981271559288604, ID: "easy"},
		{String: "hello colorhash", Value: 4155814819593785823, ID: "hello"}}
	for _, tc := range testStrings {
		t.Run(tc.ID, func(t *testing.T) {
			hash := HashString(tc.String)
			if hash != tc.Value {
				t.Errorf("%s :: Hash resulted in value %d, but expected value is %d", tc.ID, hash, tc.Value)
			}
		})
	}
}
