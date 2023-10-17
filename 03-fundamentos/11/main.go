package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	jeferson := Cliente{
		Nome:  "Jeferson",
		Idade: 29,
		Ativo: true,
	}
	jeferson.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", jeferson.Nome, jeferson.Idade, jeferson.Ativo)

}
