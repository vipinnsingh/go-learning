package main

import (
	"fmt"
	"sync"
)

func a(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println("function A is running and i : ", i)
	}

}

func main() {

	wg := sync.WaitGroup{}

	wg.Add(3)

	go c(&wg)
	go a(&wg)

	go b(&wg)
	wg.Wait()
	fmt.Println("*************")
	fmt.Println("-----------------")

}

func b(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println("function B is running and i : ", i)
	}

}

func c(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println("function C is running and i : ", i)
	}

}
