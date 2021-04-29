package main

import (
	"flag"
	"fmt"

	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGoup sync.WaitGroup

	fmt.Println("Start")

	for i := 0; i < count; i++ {
		waitGoup.Add(1)
		go func(x int) {
			defer waitGoup.Done()
			fmt.Print(x, " ")
		}(i)
	}

	waitGoup.Wait()
	fmt.Print("\nEnd")
}
