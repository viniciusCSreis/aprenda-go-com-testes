package main

import "testing"

func TestSoma(t *testing.T) {
	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperando := 6
		if esperando != resultado {
			t.Errorf("resultado %d, esperando %d, dado %v", resultado, esperando, numeros)
		}
	})

}
