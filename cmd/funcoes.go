package main

import "fmt"

func main() {
    x, y := 10, 5
    soma, subtracao := calculos(x, y)
    fmt.Println("Soma:", soma, "Subtração:", subtracao)

	cidade := "Maceió"
	temperaturas := []float64{24.6, 26.0, 29.3}

	media := calcularMedia(temperaturas...)
	fmt.Printf("Média em %s: %.1f°C\n", cidade, media)
}

func soma(a int, b int) int {  
    return a + b
}

func calculos(x, y int) (int, int) {
    soma := x + y
    subtracao := x - y
    return soma, subtracao
}

func calcularMedia(temps ...float64) float64 {
	total := 0.0
	for _, temp := range temps {
		total += temp
	}
	return total / float64(len(temps))
}