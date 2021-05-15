package cache

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	lru := Constructor(4)
	cacheAccessString := []string{"A", "B", "C", "D", "E", "D", "F"}
	for _, char := range cacheAccessString {
		if value := lru.Get(char); value == "none" {
			lru.Put(char, "value"+char)
		}
	}

	var result []string
	for e := lru.list.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(Pair).key)
	}

	var wantResult = []string{"F", "D", "E", "C"}
	for key, char := range result {
		if char != wantResult[key] {
			t.Fatalf("lur = %v, want %v", result, wantResult)
		}
	}

	for e := lru.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%+v\n", e.Value)
	}
}
