package main

import "fmt"

func main() {
	slices()
}

func arrays() {
	var numeros [3]int
	numeros[0] = 10
	numeros[1] = 20

	nomes := [2]string{"Maria", "Luís"}

	fmt.Println(numeros)
	fmt.Println(nomes[1])
}

func slices() {
	frutas := []string{"Maça", "Banana"}

	frutas = append(frutas, "Laranja")

	for i, fruta := range frutas {
		fmt.Printf("%d: %s\n", i, fruta)
	}
}