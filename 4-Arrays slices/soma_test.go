package main

import "testing"

func TestSoma(t *testing.T) {

	t.Run("colecao de 5 numeros", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, numeros)
		}
	})
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		final := numeros[1:]
		somas = append(somas, Soma(final))
	}
	return somas
}
