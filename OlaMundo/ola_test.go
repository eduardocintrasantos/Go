package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Ola, Mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
