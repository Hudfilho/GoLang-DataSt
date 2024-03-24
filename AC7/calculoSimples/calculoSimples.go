package main

import "fmt"

func main() {
	var n1, n2, v1, v2 float32
	var cod1, cod2 int

	fmt.Scan(&cod1, &n1, &v1)
	fmt.Scan(&cod2, &n2, &v2)

	total := n1*v1 + n2*v2

	fmt.Printf("VALOR A PAGAR: R$ %.2f \n", total)
}
