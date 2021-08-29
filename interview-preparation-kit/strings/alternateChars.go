package main

import (
	"fmt"
)

func alternatingCharacters(s string) int32 {
	c := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			c++
		}
	}
	return int32(c)

}

func main() {
	fmt.Println(alternatingCharacters("AAAA"))
}
