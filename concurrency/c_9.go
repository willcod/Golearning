package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func introduction() {
	fmt.Println("This code shows how to use 'select' statement")
}

func gen(min, max int, numCh chan int, end chan bool) {
	for {
		select {
		case numCh <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime out")
		}
	}
}

func main() {
	introduction()

	rand.Seed(time.Now().UnixNano())
	numCh := make(chan int)
	end := make(chan bool)

	if len(os.Args) < 2 {
		fmt.Println("Please input a number")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	go gen(0, 2*n, numCh, end)

	for i := 0; i < n; i++ {
		fmt.Print(<-numCh, " ")
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	end <- true
}
