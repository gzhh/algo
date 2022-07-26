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
			t.Fatalf("lur list = %v, want %v", result, wantResult)
		}
	}

	wantResult = []string{"valueF", "valueD", "valueE", "valueC"}
	for key, char := range result {
		if e, ok := lru.cache[char]; !ok || e.Value.(Pair).value != wantResult[key] {
			t.Fatalf("lur cache element value = %v, want %v", e.Value.(Pair).value, wantResult[key])
		}
	}

	for e := lru.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%+v\n", e.Value)
	}

	for key, element := range lru.cache {
		fmt.Printf("%s => %+v\n", key, element.Value)
	}
}
