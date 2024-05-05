package main

/*
Chamaram todos os amigos e propuseram o seguinte: com as figurinhas em mãos,
cada um tentava fazer uma troca com o amigo que estava mais perto seguindo a seguinte regra:
cada um contava quantas figurinhas tinha. Em seguida, eles tinham que dividir as figurinhas de cada um em pilhas do mesmo tamanho,
no maior tamanho que fosse possível para ambos. Então, cada um escolhia uma das pilhas de figurinhas do amigo para receber.
Por exemplo, se Ricardo e Vicente fossem trocar as figurinhas e tivessem respectivamente 8 e 12 figuras,
ambos dividiam todas as suas figuras em pilhas de 4 figuras (Ricardo teria 2 pilhas e Vicente teria 3 pilhas)
e ambos escolhiam uma pilha do amigo para receber.
*/

import "fmt"

func mdc(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var N, F1, F2 int
	fmt.Scanln(&N)

	for N != 0 {
		fmt.Scanln(&F1, &F2)
		fmt.Println(mdc(F1, F2))
		N--
	}
}
