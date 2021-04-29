package main

import (
	"fmt"
	"sync"
)

func writeToChannel(c chan int, x int, wg *sync.WaitGroup) {
	fmt.Println(x)

	c <- x
	close(c)
	defer wg.Done()

	fmt.Println(x)
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go writeToChannel(c, 10, &wg)

	fmt.Println("Read:", <-c)

	wg.Wait()

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
