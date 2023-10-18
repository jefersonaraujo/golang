package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)

}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}
func Desativavao(pessoa Pessoa) {
	pessoa.Desativar()
}
func main() {
	jeferson := Cliente{
		Nome:  "Jeferson",
		Idade: 29,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}
	Desativavao(minhaEmpresa)
	//jeferson.Desativar()

}
