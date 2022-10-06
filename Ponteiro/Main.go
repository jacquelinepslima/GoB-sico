package main

import (
	"fmt"
)

func main() {
	//var saldo float64
	var saldo = new(float64)
	fmt.Print("Digite o seu saldo: ")
	//fmt.Scanf("%f \n", &saldo)
	fmt.Scanf("%f \n", saldo)
	//alcularRendimento(&saldo)
	calcularRendimento(saldo)
	//fmt.Printf("Seu saldo com rendimentos é de R$ %.2f \n", saldo)
	fmt.Printf("Seu saldo com rendimentos é de R$ %.2f \n", *saldo)
}

func calcularRendimento(saldo *float64) {
	*saldo += *saldo * 0.03
}
