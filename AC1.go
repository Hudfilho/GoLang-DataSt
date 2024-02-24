package main

import "fmt"

//Exercicio 1: funcao calcula media
func calculaMedia(nParam int) float32 {
	var soma, item float32
	for i := 0; i < nParam; i++ {
		fmt.Println("Informe um numero: ")
		fmt.Scanln(&item)
		soma += item
	}
	return soma / float32(nParam)
}

//Exercicio 2: funcao verifica paridade
func verificaParidade(num int) {
	if (num % 2) == 0 {
		fmt.Println("numero par")
	} else {
		fmt.Println("numero impar")
	}
}

//Exercicio 3: funcao minha Potencia
func minhaPotencia(base, expoente int) int {
	resultado := base
	if expoente == 0 {
		return 1
	} else {
		for i := 1; i < expoente; i++ {
			resultado *= base
		}
		return resultado
	}
}

//Exercicio 4: funcao converter Celsius->Fahrenheit
func converteCelsiusParaFahrenheit(celsius int) float32 {
	return float32(celsius)*9/5 + 32
}

func main() {
	fmt.Println(calculaMedia(3)) //deve requisitar 3 números ao usuario e retornar a media aritm
	//verificaParidade(16) //deve printar "é par"
	//verificaParidade(11) //deve printar "é impar"
	//fmt.Println(minhaPotencia(10, 0)) //deve retornar 1
	//fmt.Println(minhaPotencia(2, 11)) //deve retornar 2048
	//fmt.Println(converteCelsiusParaFahrenheit(32)) //deve retornar 89.6
}
