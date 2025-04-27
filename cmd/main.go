package main

import "fmt"

func main() {
	var nome string = "Gabriel"
	var idade int = 21

	var (
		peso float64 = 72.5
		altura int = 184
	)

	fmt.Printf("%s tem %d anos, pesa %.1fkg e mede %dcm\n", nome, idade, peso, altura)

}
