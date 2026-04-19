package fuzz

import (
	"regexp"
	"testing"
)

func FuzzRegex(f *testing.F) {
	f.Add("image_data_1_timestamp_123")

	f.Fuzz(func(t *testing.T, input string) {
		_, err := regexp.MatchString(`^image_data_\d+_timestamp_\d+$`, input)

		if err != nil {
			t.Errorf("Regex failed: %v", err)
		}
	})
}
