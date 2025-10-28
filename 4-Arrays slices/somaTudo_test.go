package main

import (
	"reflect"
	"testing"
)

func TestSomaTudo(t *testing.T) {
	t.Run("Teste soma tudo", func(t *testing.T) {
		resultado := SomaTudo([]int{1, 2}, []int{0, 9})
		esperado := []int{3, 9}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado '%d', esperado '%d'", resultado, esperado)
		}
	})
}

func TestSomaTodoOResto(t *testing.T) {
	resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
	esperado := []int{2, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}
