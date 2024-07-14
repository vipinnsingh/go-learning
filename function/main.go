package main

import "fmt"

func main() {

	// verdic functions

	res := add(3, 45, 245, 34, 2345, 24)
	fmt.Printf("res: %v\n", res)

	// anonymous function

	func(num *int) int {
		*num = *num + 1
		return *num
	}(&res)
	fmt.Printf("res: %v\n", res)
}

func add(number ...int) int {

	var sum int
	for _, val := range number {
		sum = sum + val
	}
	return sum
}
