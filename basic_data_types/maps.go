package main

import "fmt"

func main() {
	m := map[string]int{
		"k1": 12,
		"k2": 13,
	}

	fmt.Println(m)

	var v = m["k1"]
	fmt.Println(v)

	v, ok := m["k3"]
	fmt.Println(ok)
	fmt.Println(v)
}
