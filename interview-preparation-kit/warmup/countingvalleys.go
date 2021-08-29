package main

import (
	"fmt"
	"strings"
)

func countingValleys(steps int32, path string) int32 {
	arr := strings.Split(path, "")
	newArr := []int{}
	for _, c := range arr {
		if string(c) == "U" {
			newArr = append(newArr, 1)
		} else {
			newArr = append(newArr, -1)
		}
	}
	fmt.Println(newArr)
	sum, count := 0, int32(0)
	for _, e := range newArr {
		sum = sum + e
		if sum == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countingValleys(8, "UDUUUDUDDD"))

}
