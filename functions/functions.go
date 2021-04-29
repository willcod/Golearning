package main

import (
	"fmt"
	"os"
	"strconv"
)

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("arguments are required")
		return
	}

	x, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	square := func(s int) int {
		return s * s
	}

	fmt.Println("The square of", x, "is", square(x))

	fmt.Println(doubleSquare(x))
}
