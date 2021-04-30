package main

import (
	"fmt"
	"time"
)

func introduction() {
	fmt.Println("This code shows how a channel of channel works")
}

var times int

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum += i
		}
		c <- sum
	case <-f:
		return
	}
}

func main() {
	introduction()

	cc := make(chan chan int)
	upper_limit := 10
	for i := 1; i < upper_limit+1; i++ {
		f := make(chan bool)
		go f1(cc, f)
		ch := <-cc
		ch <- i
		for sum := range ch {
			fmt.Print("Sum(", i, ")=", sum)
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}
