package main

import (
	"fmt"
)

func main() {
	var n int
	var l, q float32

	fmt.Scan(&n, &l, &q)

	nomes := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nomes[i])
	}

	i := 0
	for (l - q) > 0 {
		l = l - q
		i++
	}
	i = i % n
	ricoMate := nomes[i]

	fmt.Print(ricoMate)
	fmt.Printf(" %.1f\n", l)
}
