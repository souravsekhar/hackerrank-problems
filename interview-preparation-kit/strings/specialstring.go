package main

import "fmt"

func checkConditions(v string) bool {
	fmt.Println(v)
	if len(v) == 1 {
		return true
	} else if len(v)%2 == 0 {
		for i := 0; i < len(v)-1; i++ {
			if v[i] == v[i+1] {
				continue
			} else {
				return false
			}
		}
	} else {
		fmt.Println("In else")
		for i := 0; i <= ((len(v)-1)/2)-1; i++ {
			fmt.Println("Inside for", string(v[i]), string(v[(len(v)-i)/2+i+1]))
			if v[i] == v[(len(v)-1)/2+i+1] {
				fmt.Println("Inside if")
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func checkSpecial(s string) int64 {
	c := 0
	for i := 0; i <= len(s)-1; i++ {
		temp := string(s[i])
		for j := i + 1; j <= len(s)-1; j++ {
			temp = temp + string(s[j])
			if checkConditions(temp) {
				c++
				fmt.Println(c)
			}
		}
	}
	fmt.Println(c + len(s))
	return int64(c + len(s))
}

func main() {
	str := "abcbaba"
	checkSpecial(str)
}
