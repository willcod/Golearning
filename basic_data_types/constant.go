package main

import "fmt"

func main() {
	const PI = 3.14
	fmt.Println(PI)

	const A float64 = 123

	const (
		zero = iota
		one
	)

	fmt.Println(one)

}
