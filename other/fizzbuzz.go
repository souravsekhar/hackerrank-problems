package main

import "fmt"

func fizzBuzz(n int32) {
	for i := 1; i <= int(n); i++ {
		if i%3 == 0 {
			if i%15 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzBuzz(15)
}
