package main

import (
	"fmt"
	"sync"
	"time"
)

func introduction() {
	fmt.Println("This code shows how to timeout a goroutine")
}

func timeout(wg *sync.WaitGroup, t time.Duration) bool {
	ch := make(chan int)
	go func() {
		defer close(ch)
		time.Sleep(5 * time.Second)

		wg.Wait()
	}()

	select {
	case <-ch:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	introduction()

	var wg sync.WaitGroup
	wg.Add(1)

	duration := time.Duration(10000) * time.Millisecond

	if timeout(&wg, duration) {
		fmt.Println("Time out")
	} else {
		fmt.Println("OK")
	}

	wg.Done()
	if timeout(&wg, duration) {
		fmt.Println("Time out")
	} else {
		fmt.Println("OK")
	}
}
