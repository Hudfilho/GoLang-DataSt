/*
Entrada
A entrada contém um inteiro N, (1 ≤ N ≤ 1000) correspondente ao número de casos de teste,
seguido de N linhas, cada linha contendo um número (1 ≤ V ≤ 10100) correspondente ao valor que João quer montar com os leds.

Saída
Para cada caso de teste, imprima uma linha contendo o número de leds que João precisa para montar o valor desejado,
seguido da palavra "leds".
*/

package main

import "fmt"

func main() {
	var N int
	var valor string

	ledDigitos := map[rune]int{
		'1': 2,
		'2': 5,
		'3': 5,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 3,
		'8': 7,
		'9': 6,
		'0': 6,
	}

	fmt.Scanln(&N)
	for N > 0 {
		fmt.Scanln(&valor)

		total := 0
		for _, i := range valor {
			total += ledDigitos[i]
		}
		fmt.Printf("%d leds\n", total)
		N--
	}
}
