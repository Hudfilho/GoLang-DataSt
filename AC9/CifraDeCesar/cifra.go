/*A entrada contém vários casos de teste.
A primeira linha de entrada contém um inteiro N que indica a quantidade de casos de teste.
Cada caso de teste é composto por duas linhas.
A primeira linha contém uma string com até 50 caracteres maiúsculos ('A'-'Z'),
que é a sentença após ela ter sido codificada através desta Cifra de César modificada.
A segunda linha contém um número que varia de 0 a 25 e que representa quantas posições cada letra foi deslocada para a direita.
*/

package main

import "fmt"

func main() {
	var n, pos int
	var cifrado string

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&cifrado)
		fmt.Scanln(&pos)
		texto := ""
		var w int
		for _, j := range cifrado {
			character := int(j)
			if (character - pos) < 'A' {
				w = 26 + character - pos
			} else {
				w = character - pos
			}
			//texto += string(w)
			texto += fmt.Sprintf("%c", w)
		}
		fmt.Println(texto)
	}
}
