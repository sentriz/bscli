package main

import "testing"

func TestSearch(t *testing.T) {
	badSince := 70
	result, err := search(1, 100, func(n int) (choice, error) {
		if n < badSince {
			return choiceGood, nil
		}
		return choiceBad, nil
	})
	if err != nil {
		t.Fatalf("error in search: %v", err)
	}
	if result != badSince {
		t.Fatalf("exp %d got %d", badSince, result)
	}
}
