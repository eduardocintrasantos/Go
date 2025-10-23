package main

import "fmt"

const olaPortugues = "Olá, "
const olaEspanhol = "Hola, "
const olaFrances = "Bonjour, "
const espanhol = "espanhol"
const frances = "frances"

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = olaFrances
	case espanhol:
		prefixo = olaEspanhol
	default:
		prefixo = olaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("", ""))
}
