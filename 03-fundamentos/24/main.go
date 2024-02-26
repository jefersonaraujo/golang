package main

import (
	"fmt"

	"go-curso/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Marca)
	fmt.Println("Resultado:", s)
	fmt.Println("Resultado soma :", s)
	fmt.Println(uuid.New())

}
