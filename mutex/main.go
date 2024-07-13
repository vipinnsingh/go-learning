package main

import (
	"fmt"
	"sync"
)

func main() {

	accountBalance := 0

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go incrementBalance(&accountBalance, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(accountBalance)

}

func incrementBalance(balance *int, wg *sync.WaitGroup, mutex *sync.Mutex) {

	defer wg.Done()

	mutex.Lock()

	*balance++

	mutex.Unlock()
}
