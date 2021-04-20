package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	a := [...]int{4, 5, 6}

	t := append(s, a[:]...)
	fmt.Println("New slice:\t", t)

	s = append(s, a[:]...)
	fmt.Println("Existing slice:\t", s)

	s = append(s, s...)
	fmt.Println("s+s:\t\t", s)
}
