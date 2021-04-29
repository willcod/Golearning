package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}

	}
	close(out)
}

func third(in <-chan int) {
	sum := 0
	for x := range in {
		sum += x
	}
	fmt.Println("The sum of the random numbers is", sum)
}

func main() {
	n := flag.Int("n", 0, "the low bound")
	m := flag.Int("m", 0, "the upper bound")
	flag.Parse()

	if *n >= *m {
		fmt.Println("n must be smaller than m")
		return
	}

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first(*n, *m, A)
	go second(B, A)
	third(B)
}
