package main

import "fmt"

func simplificador(x, y int) (int, int) {
	a, b := x, y
	for y != 0 {
		x, y = y, x%y
	}
	mdc := x
	if x < 0 {
		mdc = -x
	}

	a /= mdc
	b /= mdc

	return a, b
}

func main() {
	var operador rune
	var repeticoes int
	var n1, n2, d1, d2 int

	fmt.Scanln(&repeticoes)

	for i := 0; i < repeticoes; i++ {

		fmt.Scanf("%d / %d %c %d / %d\n", &n1, &d1, &operador, &n2, &d2)

		x := 0
		y := 0
		switch operador {
		case '+':
			x = (n1*d2 + n2*d1)
			y = (d1 * d2)
		case '-':
			x = (n1*d2 - n2*d1)
			y = (d1 * d2)
		case '*':
			x = (n1 * n2)
			y = (d1 * d2)
		case '/':
			x = (n1 * d2)
			y = (n2 * d1)
		}
		a, b := simplificador(x, y)
		fmt.Printf("%d/%d = %d/%d\n", x, y, a, b)
	}
}
