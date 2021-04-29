package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	numOfRoutines := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()

	fmt.Println(*numOfRoutines, "routines will be created")

	for i := 0; i < *numOfRoutines; i++ {
		go func(x int) {
			fmt.Print(x, " ")
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("\nEnd")
}
