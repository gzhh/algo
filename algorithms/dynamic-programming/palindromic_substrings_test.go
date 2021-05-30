package dynamicprogramming

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	var stringArray = []string{
		"abc",
		"aaa",
	}
	var wants = []int{3, 6}
	for i, want := range wants {
		if got := CountSubstrings(stringArray[i]); got != want {
			t.Fatalf("CountSubstrings(s, %v) = %v, want %v", stringArray[i], got, want)
		}
	}
}
