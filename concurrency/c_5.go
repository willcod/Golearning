package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go Print(ch, &wg)

	go func() {
		for i := 1; i <= 11; i++ {
			ch <- i
		}
		close(ch)
		defer wg.Done()
	}()

	wg.Wait()

	_, ok := <-ch
	if ok {
		fmt.Println("\nChannel is open!")
	} else {
		fmt.Println("\nChannel is closed!")
	}
}

func Print(ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		fmt.Print(n, " ")
	}
	defer wg.Done()
}
