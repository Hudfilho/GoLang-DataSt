package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio  *No
	topo    *No
}

func (f *Fila) percorre() {
	if f.inicio == nil {
		fmt.Println("Fila vazia")
	} else {
		no := f.inicio
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}

}

func (f *Fila) insere(valor int) {
	novoNo := &No{valor: valor}
	if f.topo == nil {
		f.inicio = novoNo
		f.topo = novoNo
	} else {
		f.topo.prox = novoNo
		f.topo = novoNo
	}
	f.tamanho++
}

func (f *Fila) remove() error {
	if f.inicio == nil {
		return fmt.Errorf("Fila vazia")
	}
	aux := f.inicio
	f.inicio = aux.prox
	aux.prox = nil

	if f.inicio == nil {
		f.topo = nil
	}
	f.tamanho--
	return nil
}

func main() {
	fila := Fila{}

	fila.insere(2)
	fila.insere(4)
	fila.insere(8)

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()
	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	err := fila.remove()
	fmt.Println(err)

	fila.insere(1)
	fila.insere(9)

	fila.percorre()
	fmt.Println(fila.tamanho)
}
