package main

import (
	"fmt"
	"sync"
)

func main() {

	// var ch1 chan string = make(chan string)
	// var ch2 chan string
	// ch1 <- "pre process"

	// go func() {
	// 	ch1 <- "sending data"
	// }()

	// fmt.Printf("ch2: %v\n", ch2)
	// fmt.Printf("ch1: %v\n", ch1)

	// fmt.Printf("ch1 data: %v\n", <-ch1)
	// fmt.Printf("ch2 data: %v\n", <-ch2)

	// buffer channels

	ch2 := make(chan int, 5)
	fmt.Printf("ch2: %v\n", ch2)
	go func() {
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
		ch2 <- 4
		ch2 <- 5
		close(ch2)
	}()

	for val := range ch2 {
		fmt.Printf("val in buffer channel ch2: %v\n", val)
	}

	var ch chan int = make(chan int)
	// fmt.Printf("ch: %v\n", ch)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go SendNumber(ch, &wg)
	go ReceiveNumber(ch, &wg)
	// time.Sleep(time.Second * 3)
	wg.Wait()

}

func SendNumber(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 11; i++ {
		ch <- i
	}
	close(ch)
}

func ReceiveNumber(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Printf("ch: %v\n", ch)
	fmt.Printf("ch data outside loop: %v\n", <-ch)
	for val := range ch {
		fmt.Printf("ch data: %v\n", val)
	}
	// close(ch)

}
