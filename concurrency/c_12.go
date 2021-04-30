package main

import (
	"fmt"
)

func introduction() {
	fmt.Println("This code shows how a buffered channel works")
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
