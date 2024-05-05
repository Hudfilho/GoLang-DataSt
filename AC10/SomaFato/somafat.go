package main

/*Leia dois valores inteiros M e N indefinidamente.
A cada leitura, calcule e escreva a soma dos fatoriais de cada um dos valores lidos.
Utilize uma variável apropriada, pois cálculo pode resultar em um valor com mais de 15 dígitos.
*/
import "fmt"

func fatorial(x int64) int64 {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func main() {
	var M, N int64
	for {
		_, err1 := fmt.Scanln(&M, &N)
		if err1 != nil {
			return
		}
		fmt.Println(fatorial(M) + fatorial(N))
	}
}
