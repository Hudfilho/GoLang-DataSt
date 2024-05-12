package main

import "fmt"

func main() {
	var N, M, P int
	fmt.Scanln(&N)
	for N != 0 {

		listaItens := make(map[string]float64)
		total := 0.0

		fmt.Scanln(&M)
		for M != 0 {
			var nomeItem string
			var precoItem float64
			fmt.Scanln(&nomeItem, &precoItem)
			listaItens[nomeItem] = precoItem
			M--
		}
		fmt.Scanln(&P)
		for P != 0 {
			var produto string
			var quantidade int
			fmt.Scanln(&produto, &quantidade)
			total += float64(quantidade) * listaItens[produto]
			P--
		}
		fmt.Printf("R$ %.2f\n", total)
		N--
	}
}
