package main

import "fmt"

func main() {
	// a := []int{1, 2, 3}
	const card = "sjv"
	fmt.Printf("card: %v\n", card)

	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("a[:1]: %v\n", a[:1])
	// fmt.Printf("a[1:]: %v\n", a[1:])

	// b := append(a[:1], 10)

	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("a again: %v\n", a)

	// c := make([]int, 3, 90)
	// fmt.Println(c)
	// d := append(c, 5, 4, 5, 67, 7, 6)
	// d[0] = 1
	// d[4] = 8
	// c = append(c, 4)
	// fmt.Println(c)

	// fmt.Println(d)
	// c := make([]int, 4)
	// c[3] = 3
	// fmt.Println(c)

	// s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// //v := []int{}
	// _ = append(s[:3], s[4:]...) // remove 3rd elem
	// fmt.Println(s)

	sports := make([]int, 5)
	sports[0] = 1
	sports[1] = 2
	sports[2] = 3
	sports[3] = 4
	sports[4] = 5

	xs := sports[1:3]
	fmt.Printf("xs: %v\n", xs)
	xs[0] = 9

	fmt.Println(xs)
	fmt.Println(sports)

	for _, v := range xs {
		fmt.Printf("Address %v And value %v\n", &v, v)
	}

}
