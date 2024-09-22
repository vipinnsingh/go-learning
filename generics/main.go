package main

import "fmt"

type Num interface {
	int | int16 | int64 | float32 | float64
}

func Add[T Num](a T, b T) T {
	return a + b
}

func Map[T Num](arr []T, mapFunc func(T) T) []T {
	for i := range arr {
		arr[i] = mapFunc(arr[i])
	}

	return arr
}
func main() {
	c := Add(1.3, 1.5)
	fmt.Printf("c: %v\n", c)

	// input : [1, 2, 3]
	// output : [2, 4, 6]
	a := []int{2, 3, 4, 5}
	arr := Map(a, func(i int) int {
		return i * 2
	})
	fmt.Printf("arr: %v\n", arr)
}
