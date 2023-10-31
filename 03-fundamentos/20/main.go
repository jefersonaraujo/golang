package main

type MyNumber int

type Number interface {
	~int | ~float64 // (~ = exceção)
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {

	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Jeferson": 1000, "Bob": 2000, "Maria": 3000}
	m2 := map[string]float64{"Jeferson": 100.20, "Bob": 2000.3, "Maria": 300.0}

	m3 := map[string]MyNumber{"Jeferson": 1000, "Bob": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))

}
