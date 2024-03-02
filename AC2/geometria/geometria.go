package geometria

//Exercicio 4: crie um pacote geometria com funcoes para calcular a area e o perimetro de um retangulo.
func Area(largura, altura float32) float32 {
	return largura * altura
}

func Perimetro(largura, altura float32) float32 {
	return (largura + altura) * 2
}
