package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Jeferson"
	e float64 = 1.2
	f ID      = 1
)

func main() {

	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 10

	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])
	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}

}
