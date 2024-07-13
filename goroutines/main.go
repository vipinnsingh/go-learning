package main

import (
	"fmt"
	"time"
)

func main() {
	// a()
	// b()
	// c()
	// d()

	go a()
	go b()
	go c()
	go d()

	time.Sleep(time.Second * 3)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("function A is running and i : ", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("function B is running and i : ", i)
	}
}

func c() {
	for i := 0; i < 10; i++ {
		fmt.Println("function C is running and i : ", i)
	}
}

func d() {
	for i := 0; i < 10; i++ {
		fmt.Println("function D is running and i : ", i)
	}
}
