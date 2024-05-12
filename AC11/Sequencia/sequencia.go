package main

import "fmt"

func main() {
	var N, atual, consecutivos int

	fmt.Scanln(&N)
	sequencia := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&sequencia[i])
	}

	for i := 0; i < N; i++ {
		if sequencia[i] != atual {
			consecutivos++
		}
		atual = sequencia[i]
	}
	fmt.Println(consecutivos)
}
