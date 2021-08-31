package main

import "fmt"

func leftRotate(arr []int, times int) []int {
	temp := 0

	for i := 0; i < times; i++ {
		temp = arr[0]
		for j := 0; j < len(arr)-1; j++ {
			arr[j] = arr[j+1]
		}
		arr[len(arr)-1] = temp
	}

	return arr

}
func main() {
	fmt.Println(leftRotate([]int{1, 2, 3, 4, 5, 6}, 3))
}
