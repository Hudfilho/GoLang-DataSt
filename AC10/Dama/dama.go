package main

/*
A entrada contém vários casos de teste. A primeira e única linha de cada caso de teste contém quatro inteiros X1, Y1, X2 e Y2
(1 ≤ X1, Y1, X2, Y2 ≤ 8).
A dama começa na casa de coordenadas (X1, Y1), e a casa de destino é a casa de coordenadas(X2, Y2).
No tabuleiro, as colunas são numeradas da esquerda para a direita de 1 a 8 e as linhas de cima para baixo também de 1 a 8.
As coordenadas de uma casa na linha X e coluna Y são (X, Y ).
O final da entrada é indicado por uma linha contendo quatro zeros.
*/

import (
	"fmt"
	"math"
)

func minMov(a, b, c, d int) int {
	if a == c && b == d {
		return 0
	} else if (a == c || b == d) || (math.Abs(float64(a-c)) == math.Abs(float64(b-d))) {
		return 1
	}
	return 2
}

func main() {
	var x1, x2, y1, y2 int

	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)

		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			return
		}
		fmt.Println(minMov(x1, y1, x2, y2))
	}

}
