package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	pingCh := make(chan string)
	pongCh := make(chan string)

	wg.Add(1)

	go ping(pingCh, pongCh, &wg)
	go pong(pingCh, pongCh, &wg)
	pingCh <- "ping"
	wg.Wait()
}

func ping(pingCh <-chan string, pongCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		<-pingCh
		fmt.Println("ping")
		pongCh <- "pong"
	}
}

func pong(pingCh chan<- string, pongCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		<-pongCh
		fmt.Println("pong")
		pingCh <- "ping"
	}
}
