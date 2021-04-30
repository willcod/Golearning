package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This code shows how to get the max number of the logical processors on the system.")
	fmt.Println("The max number of processors:", runtime.GOMAXPROCS(0))
}
