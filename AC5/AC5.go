package main

import "fmt"

//exercicio 2: Tipos de triangulos
func tiposTriangulos() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}

	switch {
	case a >= (b + c):
		fmt.Println("NAO FORMA TRIANGULO")
	case a*a == (b*b + c*c):
		fmt.Println("TRIANGULO RETANGULO")
	case a*a > (b*b + c*c):
		fmt.Println("TRIANGULO OBTUSANGULO")
	case a*a < (b*b + c*c):
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	if a == b && a == c {
		fmt.Println("TRIANGULO EQUILATERO")
	} else if a == b || a == c || b == c {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}

//exercicio 3: Crescimento populacional
func crescimento() {
	var T, PA, PB int
	var G1, G2 float32
	fmt.Scanln(&T)

	for T != 0 {
		fmt.Scanln(&PA, &PB, &G1, &G2)
		i := 0
		for PA <= PB {
			PA = int(float32(PA) * (1 + G1/100))
			PB = int(float32(PB) * (1 + G2/100))
			i++
			if i > 100 {
				fmt.Println("Mais de 1 seculo.")
				break
			}
		}

		if i <= 100 {
			fmt.Println(i, "anos.")
		}
		T--
	}
}

func main() {
	//exercicio 1: hello world
	fmt.Println("Hello World!")

	tiposTriangulos()

	crescimento()
}
