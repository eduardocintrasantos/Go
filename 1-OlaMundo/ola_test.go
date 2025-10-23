package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Eduardo", "")
		esperado := "Olá, Eduardo"

		verificaMensagemCorreta(t, resultado, esperado)
	})
	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})
	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Eduardo", "espanhol")
		esperado := "Hola, Eduardo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
	t.Run("em frances", func(t *testing.T) {
		resultado := Ola("Eduardo", "frances")
		esperado := "Bonjour, Eduardo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
