package main

import "fmt"

func main() {
	oneD := [4]int{1, 2, 4, 3}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	triD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}

	fmt.Println("The length of", oneD, "is", len(oneD))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The lenght of", triD, "is", len(triD))

	for i := 0; i < len(triD); i++ {
		v := triD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
			fmt.Println()
		}
	}

	for _, v := range triD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
			fmt.Println()
		}
	}
}
