package main

import "fmt"

var pos int

func repeatedString(s string, n int64) int64 {
	count := n / int64(len(s))
	r := int(n % int64(len(s)))
	occ := int64(0)
	fmt.Println(count, r)

	for _, c := range s {
		if string(c) == "a" {
			occ++
		}
	}
	fmt.Println(occ)
	total := occ * count
	for i := 0; i < r; i++ {
		if string(s[i]) == "a" {
			total++
		}
	}
	return total

}

func main() {
	fmt.Println(repeatedString("aba", 10))
}
