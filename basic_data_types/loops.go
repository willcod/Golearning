package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()

	i := 10
	for {
		if i < 0 {
			break
		}

		fmt.Print(i, " ")
		i--
	}

	i = 0
	flag := true
	for ok := true; ok; ok = flag {
		if i > 10 {
			flag = false
		}

		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	a := [5]int{0, 1, -1, 2, -1}
	for i, value := range a {
		fmt.Println("index: ", i, "value: ", value)
	}

}
