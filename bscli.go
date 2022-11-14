package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <start> <end>\n", os.Args[0])
		os.Exit(1)
	}

	start, _ := strconv.Atoi(os.Args[1])
	end, _ := strconv.Atoi(os.Args[2])

	result, err := search(start, end, func(n int) (choice, error) {
		fmt.Printf("%d?\n", n)
		var input string
		if _, err := fmt.Scanf("%s", &input); err != nil {
			return choiceNone, fmt.Errorf("scan: %w", err)
		}
		choice, err := parseChoice(input)
		if err != nil {
			return choiceNone, fmt.Errorf("parse: %w", err)
		}
		return choice, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stdout, "error searching: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func search(start int, end int, check func(n int) (choice, error)) (int, error) {
	var mid int
	for start <= end {
		mid = (start + end) >> 1
		result, err := check(mid)
		if err != nil {
			return 0, fmt.Errorf("check: %w", err)
		}
		switch result {
		case choiceGood:
			start = mid + 1
		case choiceBad:
			end = mid - 1
		}
	}
	return mid, nil
}

type choice int

const (
	choiceNone choice = iota
	choiceGood
	choiceBad
)

func parseChoice(str string) (choice, error) {
	switch str := strings.ToLower(str); str {
	case "b", "bad", "new":
		return choiceBad, nil
	case "g", "good", "old":
		return choiceGood, nil
	default:
		return choiceNone, fmt.Errorf("unknown choice")
	}
}
