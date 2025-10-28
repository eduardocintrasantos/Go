package main

import "testing"

func TestArea(t *testing.T) {
	t.Run("retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.00, 6.00}
		resultado := retangulo.Area()
		esperado := 72.00

		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
		}
	})

	t.Run("circulo", func(t *testing.T) {
		circulo := Circulo{10}
		resultado := circulo.Area()
		esperado := 314.1592653589793

		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
		}
	})
}
