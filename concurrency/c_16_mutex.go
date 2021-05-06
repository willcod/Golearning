package main

import (
	"fmt"
	"sync"
	"time"
)

func introduction() {
	fmt.Println("This code shows how to use mutex")
}

var (
	m sync.Mutex
	g int
)

func change(num int) {
	m.Lock()
	time.Sleep(time.Second)
	g++
	if g%10 == 0 {
		g = g - num*10
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	num := g
	m.Unlock()
	return num
}

func main() {
	introduction()

	var wg sync.WaitGroup
	fmt.Print(read(), " ")
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			change(num)
			fmt.Print("-> ", read())
		}(i)
	}
	wg.Wait()
	fmt.Println("-> ", read())
}
