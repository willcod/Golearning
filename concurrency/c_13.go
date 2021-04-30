package main

import (
	"fmt"
	"math/rand"
	"time"
)

func introduction() {
	fmt.Println("This code shows how a nil channel works")
}

func add(c chan int) {
	sum := 0
	t := time.NewTimer(2 * time.Second)

	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	introduction()

	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(3 * time.Second)
}
