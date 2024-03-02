package main

import (
	"fmt"
	"math"
	g "projeto/geometria"
)

// Exercicio 2: metodo que recebe string e retorne uma nova string, com a sequencia invertida. (pesquisar Runes e String)
func inverteString(s string) string {
	inverte := []rune(s)

	for i, j := len(s)-1, 0; i > len(s)/2-1; i, j = i-1, j+1 {
		inverte[i], inverte[j] = inverte[j], inverte[i]
	}
	return string(inverte)
}

//Exercicio 3: crie um ponto em um espaco cartesiano (X,Y) e calcule a distancia do ponto ate a origem.

type Ponto struct {
	X float64
	Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	//raiz quadrada de x^2+y^2
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {

	//Exercicio 1: Vetor de int com 10 elementos -> inicialize o vetor e imprima os valores.
	var vetor [10]int
	vetor[0] = 17
	vetor[1] = 2
	vetor[2] = 356
	vetor[3] = 4
	vetor[4] = 54
	vetor[5] = 67
	vetor[6] = 73
	vetor[7] = 82
	vetor[8] = 91
	vetor[9] = 101
	for i := 0; i < len(vetor); i++ {
		fmt.Println(vetor[i])
	}

	ponto := Ponto{X: 3, Y: 4}
	fmt.Println(ponto.DistanciaOrigem())

	//Exercicio 2:
	var palavra string
	fmt.Println("Informe uma string para ser invertida: ")
	fmt.Scanln(&palavra)
	fmt.Println(inverteString(palavra))

	//Exercicio 4: crie um pacote geometria com funcoes para calcular a area e o perimetro de um retangulo.
	fmt.Println(g.Area(12, 78))
	fmt.Println(g.Perimetro(5, 6))

}
