package main

import "fmt"

const palavraOla = "Olá, "

func Ola(nome string) string {
	if nome != "" {
		return palavraOla + nome
	} else {
		return palavraOla + "mundo"
	}

}
func main() {
	fmt.Println(Ola(""))
}
