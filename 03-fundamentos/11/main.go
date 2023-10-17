package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	jeferson := Cliente{
		Nome:  "Jeferson",
		Idade: 29,
		Ativo: true,
	}
	jeferson.Ativo = false
	jeferson.Endereco.Cidade = "Fortaleza"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", jeferson.Nome, jeferson.Idade, jeferson.Ativo)

}
