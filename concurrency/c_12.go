package main

import (
	"fmt"
	"sync"
	"time"
)

func introduction() {
	fmt.Println("This code shows how a buffered channel works")
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

	nums := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case nums <- i:
		default:
			fmt.Println("No enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-nums:
			fmt.Println(num)
		default:
			fmt.Println("No more to be done")
			break
		}
	}
}
