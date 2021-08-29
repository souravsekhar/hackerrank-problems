package main

import (
	"fmt"
)

func isValid(s string) string {
	m := map[string]int{}
	for _, v := range s {
		if _, ok := m[string(v)]; !ok {
			m[string(v)] = 1
		} else {
			m[string(v)] = m[string(v)] + 1
		}
	}
	largest := 0
	smallest := len(s)
	for key := range m {
		fmt.Println(key)
		if m[key] > largest {
			largest = m[key]
		} else if m[key] < smallest {
			smallest = m[key]
		}
	}
	fmt.Println(largest, smallest)
	if largest-smallest > 1 {
		return "NO"
	}
	return "YES"
}

func main() {
	fmt.Println(isValid("abcdefghhgfedecba"))
}
