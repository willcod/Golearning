package main

import (
	"fmt"
	"time"
)

func introduction() {
	fmt.Println("This code shows how to specify the running order of goroutines")
}

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C()")
}

func main() {
	introduction()

	a := make(chan struct{})
	b := make(chan struct{})
	c := make(chan struct{})

	go C(c)
	go A(a, b)
	go C(c)
	go B(b, c)
	go C(c)

	close(a)
	time.Sleep(3 * time.Second)
}
