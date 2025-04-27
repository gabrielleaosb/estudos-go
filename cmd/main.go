package main

import "fmt"

func main() {
	fmt.Println("Salve!!!!")

	variaveis()
	funcoes()
}

func variaveis() {
	var nome string = "Go"
	versao := 1.24
	fmt.Println("Aprendendo", nome, "(versão", versao, ")") 
}

func funcoes() {
	resultado := soma(5, 3)
	fmt.Println("5 + 3 =", resultado)
}

func soma(a, b int) int {
	return a + b
}