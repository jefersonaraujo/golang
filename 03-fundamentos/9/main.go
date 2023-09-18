package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 23, 4, 41, 1, 4, 14, 141, 1, 41, 41, 41, 4141))

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}
