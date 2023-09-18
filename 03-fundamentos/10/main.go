package main

import (
	"fmt"
)

func main() {
	total := func() int {

		return sum(1, 23, 4, 41, 1, 4, 14, 141, 1, 41, 41, 41, 4141) * 2
	}()

	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}
