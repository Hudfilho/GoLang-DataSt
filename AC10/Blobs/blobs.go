package main

/*
No planeta Alpha vive a criatura Blobs, que come precisamente 1/2 de seu suprimento de comida disponível todos os dias.
Escreva um algoritmo que leia a capacidade inicial de suprimento de comida (em Kg),
e calcule quantos dias passarão antes que Blobs coma todo esse suprimento até restar um quilo ou menos.
*/

import "fmt"

func diasRestantes(x float32) int {
	dias := 0
	for x > 1 {
		x /= 2
		dias++
	}
	return dias
}

func main() {
	var N int
	var C float32
	fmt.Scanln(&N)

	for N != 0 {
		fmt.Scanln(&C)
		fmt.Printf("%d dias\n", diasRestantes(C))
		N--
	}
}
