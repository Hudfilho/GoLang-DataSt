/*A entrada consiste de diversos casos de teste, cada um em uma linha.
Cada linha contém dois inteiros D e N (1 ≤ D ≤ 9, 1 ≤ N < 10100 ),
representando, respectivamente, o dígito que está apresentando problema na máquina e o número que foi negociado
originalmente no contrato (que podem ser grande, pois Modernolândia tem sido acometida por hiperinflação nas últimas décadas).
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var D int
	var N string
	for {
		fmt.Scan(&D)
		fmt.Scanln(&N)
		if D == 0 && N == "0" {
			break
		}

		V := strings.ReplaceAll(N, fmt.Sprintf("%d", D), "")

		V = strings.TrimLeft(V, "0")

		if V == "" {
			V = "0"
		}
		fmt.Println(V)
	}
}
