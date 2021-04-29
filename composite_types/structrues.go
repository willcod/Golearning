package main

import "fmt"

type Person struct {
	Name   string
	Height int
	Weight int
}

func main() {
	var p Person
	fmt.Println(p)
}
