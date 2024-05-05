package main

/*

A entrada contém apenas 1 caso de teste. A primeira linha de entrada contem um único inteiro N,
que indica a quantidade de valores que serão lidos para X (1 ≤ X ≤ 2000)
logo em seguida. Com certeza cada número não aparecerá mais do que 20 vezes na entrada de dados.
*/

import (
	"fmt"
	"sort"
)

func mapearFreq(x []int) {
	frequencia := make(map[int]int)

	for _, n := range x {
		frequencia[n]++
	}

	var numeros []int
	for num := range frequencia {
		numeros = append(numeros, num)
	}
	sort.Ints(numeros)

	for _, num := range numeros {
		fmt.Printf("%d aparece %d vez(es)\n", num, frequencia[num])
	}
}

func main() {
	var N int
	fmt.Scanln(&N)

	entradas := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&entradas[i])
	}
	mapearFreq(entradas)
}
