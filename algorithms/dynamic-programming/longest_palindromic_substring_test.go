package dynamicprogramming

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var stringArray = []string{
		"babad",
		"cbbd",
		"a",
		"ac",
	}
	var wants = []string{
		"bab",
		"bb",
		"a",
		"a",
	}
	for i, want := range wants {
		if got := LongestPalindrome(stringArray[i]); got != want {
			t.Fatalf("LongestPalindrome(s, %v) = %v, want %v", stringArray[i], got, want)
		}
	}
}
