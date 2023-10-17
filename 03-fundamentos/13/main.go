package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)

}
func main() {
	jeferson := Cliente{
		Nome:  "Jeferson",
		Idade: 29,
		Ativo: true,
	}
	jeferson.Desativar()

	//fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", jeferson.Nome, jeferson.Idade, jeferson.Ativo)

}
