package main

import "fmt"

func square(n *int) {
	*n = *n * *n
}

func main() {
	i := -10
	j := -25

	pI := &i
	pJ := &j

	fmt.Println("PI memory:", pI, "PI value:", *pI)
	fmt.Println("PJ memory:", pJ, "PJ value:", *pJ)

	*pI--
	fmt.Println("i:", i)

	square(&i)
	fmt.Println("Square i:", i)
}
