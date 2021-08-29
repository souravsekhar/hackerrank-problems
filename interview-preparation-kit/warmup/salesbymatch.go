package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	m := map[int32]int32{}
	pairs := int32(0)
	for _, v := range ar {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}
	fmt.Println(m)

	for key := range m {
		pairs = pairs + m[key]/2
	}
	return pairs
}

func main() {
	input := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(9, input))
}
