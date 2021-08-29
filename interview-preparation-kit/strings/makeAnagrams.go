package main

import (
	"fmt"
	"math"
)

func makeAnagram(a string, b string) int32 {
	// c := int32(0)
	// m := map[string]string{}
	// for _, v := range a {
	// 	if strings.Contains(b, string(v)) {
	// 		// if _, found := m[string(v)]; !found {
	// 		// 	fmt.Println(string(v))
	// 		// 	m[string(v)] = string(v)
	// 		continue
	// 	} else {
	// 		c++
	// 		a = strings.ReplaceAll(a, string(v), "")
	// 	}
	// }
	// fmt.Println(m)
	// for _, v := range b {
	// 	if strings.Contains(a, string(v)) {
	// 		continue
	// 	} else {
	// 		c++
	// 		b = strings.ReplaceAll(b, string(v), "")
	// 	}
	// }
	m := map[string]int{}
	total := 0
	for _, v := range a {
		if _, ok := m[string(v)]; !ok {
			m[string(v)] = 1
		} else {
			m[string(v)] = m[string(v)] + 1
		}
	}
	for _, v := range b {
		if _, ok := m[string(v)]; !ok {
			m[string(v)] = 1
		} else {
			m[string(v)] = m[string(v)] - 1
		}
	}
	for key := range m {
		total = total + int(math.Abs(float64(m[key])))
	}
	return int32(total)

}

func main() {
	a := "fcrxzwscanmligyxyvym"
	b := "jxwtrhvujlmrpdoqbisbwhmgpmeoke"
	r := makeAnagram(a, b)
	fmt.Println(r)
}
