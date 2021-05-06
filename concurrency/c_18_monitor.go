package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func introduction() {
	fmt.Println("This code shows how to use monitor goroutine")
}

var (
	readValue  = make(chan int)
	writeValue = make(chan int)
)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Print(value, " ")
		case readValue <- value:
		}
	}
}

func main() {
	introduction()
	n := 10
	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}
	wg.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}
