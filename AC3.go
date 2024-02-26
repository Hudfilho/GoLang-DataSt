package main

import "fmt"

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) MudaEmail(newEmail string) {
	c.Email = newEmail
}

func main() {
	var contatos [5]Contato
	for {
		var interacao int
		var nome, email string
		fmt.Println("Informe o procedimento!: (1)Adicionar Contato (2)Excluir Contato (3)Sair")
		fmt.Scanln(&interacao)
		switch interacao {
		case 1:
			fmt.Println("Informe o Nome")
			fmt.Scanln(&nome)
			fmt.Println("Informe o Email")
			fmt.Scanln(&email)
			contato := Contato{Nome: nome, Email: email}
			adicionarContato(contato, &contatos)
			fmt.Println(contatos)
		case 2:
			excluirContato(&contatos)
			fmt.Println(contatos)
		case 3:
			return

		}
	}
}

func adicionarContato(novoContato Contato, contatos *[5]Contato) {
	for i, contato := range contatos {
		if (contato == Contato{}) {
			contatos[i] = novoContato
			return
		}
	}
	fmt.Println("lista cheia!")
}
func excluirContato(contatos *[5]Contato) {
	for i := len(contatos) - 1; i >= 0; i-- {
		if (contatos[i] != Contato{}) {
			contatos[i] = Contato{}
			return
		}
	}
}
