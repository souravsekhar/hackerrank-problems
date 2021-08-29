package main

import "fmt"

func fibonacci(n int) {
	start, next, temp := 1, 1, 0
	for i := 0; i < n; i++ {
		temp = next
		next = start + next
		start = temp
		fmt.Println(next)
	}
}

func integerPalindrome(n int) {
	r, rev, temp := 0, 0, n
	for temp > 0 {
		r = temp % 10
		rev = rev*10 + r
		temp = temp / 10
	}

	if rev == n {
		fmt.Println("Integer palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}

func secondlargest(arr []int) {
	largest, secondlargest := arr[0], arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			secondlargest = largest
			largest = arr[i]
		} else if arr[i] > secondlargest {
			secondlargest = arr[i]
		}
	}

	fmt.Println(largest, secondlargest)
}

func stringPalindrome(str string) {
	var rev string
	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + string(str[i])
	}
	if rev == str {
		fmt.Println("palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}

func checkPrimeNumbers(n int) bool {
	if n <= 1 {
		return false
	} else {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}

func findPrimeNumbers(nrange int) {
	for i := 0; i <= nrange; i++ {
		if checkPrimeNumbers(i) {
			fmt.Println(i)
		}
	}
}

func leftRotateArray(arr []int, times int) {
	for i := 0; i < times; i++ {
		temp := arr[0]
		for j := 0; j < len(arr)-1; j++ {
			arr[j] = arr[j+1]
		}
		arr[len(arr)-1] = temp
	}

	fmt.Println(arr)
}

func rightRotateArray(arr []int, times int) {
	for i := 0; i < times; i++ {
		temp := arr[len(arr)-1]
		for j := len(arr) - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[0] = temp
	}
	fmt.Println(arr)
}

func swapNumbers(a int, b int) {
	fmt.Println("Before", a, b)
	if a != b {
		a = a * b
		b = a / b
		a = a / b
	}
	fmt.Println("After", a, b)
}

func sortArray(arr []int) {
	var temp int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}

// TODO: Permutation and combinations
// Find pairs in array with sum = input
// anagram
// reverse a string with reversed words
func main() {
	// fibonacci(10)
	// integerPalindrome(1221)
	// secondlargest([]int{1, 2, 3, 1, 4, 5, 10, 9, 12, 0})
	// stringPalindrome("sts")
	// checkPrimeNumbers(12)
	// findPrimeNumbers(100)
	// leftRotateArray([]int{1, 2, 3, 4, 5}, 1)
	// rightRotateArray([]int{1, 2, 3, 4, 5}, 3)
	// swapNumbers(10, 20)
	sortArray([]int{3, 5, 1, 7, 23, 2, 12})

}
