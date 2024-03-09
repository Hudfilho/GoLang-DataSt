package main

import "fmt"

//Solução recursiva para o problema da Torre de Hanói.
func hanoi(n int, origem, trabalho, destino string) {
	if n == 0 {
		return
	}
	hanoi(n-1, origem, destino, trabalho)
	fmt.Printf("O anel %d foi movido de %s para %s\n", n, origem, destino)
	hanoi(n-1, trabalho, origem, destino)
}

func parON(array []int, x int) []int {
	n1, n2 := 0, len(array)-1

	for n1 < n2 {
		soma := array[n1] + array[n2]
		if soma == x {
			return []int{array[n1], array[n2]}
		} else if soma > x {
			n2--
		} else {
			n1++
		}
	}
	return []int{-1, -1}
}

func main() {
	//hanoi(1, "origem", "trabalho", "destino")
	array := []int{1, 4, 6, 8, 11, 15, 23}
	fmt.Println(parON(array, 24))
}
