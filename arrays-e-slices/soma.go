package main

func Soma(dados []int) int {
	soma := 0
	for _, numero := range dados{
		soma += numero
	}
	return soma
}
