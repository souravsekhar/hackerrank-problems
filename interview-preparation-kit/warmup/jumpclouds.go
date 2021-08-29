package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	count := 0
	for i := 0; i < len(c)-1; i += 2 {
		if c[i] == 1 {
			count++
			i--
		} else {
			count++
		}

	}
	fmt.Println(count)
	return int32(count)
}

func main() {
	// input := []int32{0, 0, 1, 0, 0, 1, 0}
	input := []int32{0, 1, 0, 0, 0, 1, 0}
	fmt.Println(jumpingOnClouds(input))
}
