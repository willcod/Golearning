package main

import (
	"fmt"
	"time"
)

func countPrint() {
	for i := 0; i < 10; i++ {
		fmt.Print("f-", i, " ")
	}
}

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print("a-", i, " ")
		}
	}()

	fmt.Println()

	go countPrint()

	time.Sleep(1 * time.Second)
}
