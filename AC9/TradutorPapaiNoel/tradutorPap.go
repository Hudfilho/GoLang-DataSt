/*
Entrada
Você irá testar o seu aplicativo com diversos nomes de paises, simulando os dados informados pelo painel de navegação do trenó.

Saída
O seu aplicativo deverá mostrar na tela a frase no idioma correto.
Caso ela não esteja cadastrada, você deverá exibir a mensagem "--- NOT FOUND ---"
para que depois dos testes você possa completar o banco de dados.
*/

package main

import "fmt"

func main() {

	mensagem := map[string]string{
		"brasil":         "Feliz Natal!",
		"alemanha":       "Frohliche Weihnachten!",
		"austria":        "Frohe Weihnacht!",
		"coreia":         "Chuk Sung Tan!",
		"espanha":        "Feliz Navidad!",
		"grecia":         "Kala Christougena!",
		"estados-unidos": "Merry Christmas!",
		"inglaterra":     "Merry Christmas!",
		"australia":      "Merry Christmas!",
		"portugal":       "Feliz Natal!",
		"suecia":         "God Jul!",
		"turquia":        "Mutlu Noeller",
		"argentina":      "Feliz Navidad!",
		"chile":          "Feliz Navidad!",
		"mexico":         "Feliz Navidad!",
		"antardida":      "Merry Christmas!",
		"canada":         "Merry Christmas!",
		"irlanda":        "Nollaig Shona Dhuit!",
		"belgica":        "Zalig Kerstfeest!",
		"italia":         "Buon Natale!",
		"libia":          "Buon Natale!",
		"siria":          "Milad Mubarak!",
		"marrocos":       "Milad Mubarak!",
		"japao":          "Merii Kurisumasu!",
	}
	for {
		pais := ""
		fmt.Scanln(&pais)
		if pais == "" {
			break
		}
		if mensagem[pais] == "" {
			fmt.Println("--- NOT FOUND ---")
		} else {
			fmt.Println(mensagem[pais])
		}
	}
}
