//Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como parâmetro e calcule a área do círculo (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi.

package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func CalcularArea(circulo circulo) float64 {
	return math.Pi * circulo.raio * circulo.raio
}

func main() {
	var raio float64
	fmt.Println("Qual o seu raio? ")
	fmt.Scan(&raio)
	c := circulo{raio}
	area := CalcularArea(c)
	fmt.Println("Sua area é: ", area)
}
