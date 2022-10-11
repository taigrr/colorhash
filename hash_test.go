package colorhash

import (
	"testing"
)

func TestHashBytes(t *testing.T) {
	testBytes := []struct{}{}
	_ = testBytes
}

func TestHashString(t *testing.T) {
	testStrings := []struct {
		String string
		Value  int
		ID     string
	}{
		{String: "", Value: 5472609002491880228, ID: "Empty string"},
		{String: "123", Value: 6449148174219763898, ID: "123"},
		{String: "it's as easy as", Value: 5908178111834329190, ID: "easy"},
		{String: "hello colorhash", Value: 893132354324239557, ID: "hello"},
	}
	for _, tc := range testStrings {
		t.Run(tc.ID, func(t *testing.T) {
			hash := HashString(tc.String)
			if hash != tc.Value {
				t.Errorf("%s :: Hash resulted in value %d, but expected value is %d", tc.ID, hash, tc.Value)
			}
		})
	}
}
