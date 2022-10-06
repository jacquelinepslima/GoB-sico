package main

import (
	"fmt"
)

func main() {
	var numero1, numero2 int
	var operacao string
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanf("%d \n", &numero1)
	fmt.Print("Digite a operação (+, -, *, /, $)")
	fmt.Scanf("%s \n", &operacao)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%d \n", &numero2)
	if operacao == "+" {
		somar(numero1, numero2)
	} else if operacao == "-" {
		resultado := subtrair(numero1, numero2)
		fmt.Printf("%d - %d = %d \n", numero1, numero2, resultado)
	} else if operacao == "*" {
		resultado := multiplicar(numero1, numero2)
		fmt.Printf("%d * %d = %d \n", numero1, numero2, resultado)
	} else if operacao == "/" {
		resultado, resto := dividir(numero1, numero2)
		fmt.Printf("Quociente = %d; Resto = %d \n", resultado, resto)
	} else if operacao == "$" {
		incremento, decremento := incrementoDecremento(numero1, numero2)
		fmt.Printf("Incremento = %d, Decremento = %d \n", incremento, decremento)
	} else {
		fmt.Println("Operação Inválida")
	}
}

//void - não retorna nada
func somar(n1 int, n2 int) {
	fmt.Printf("%d + %d = %d \n", n1, n2, n1+n2)
}

//retorna algo
func subtrair(n1 int, n2 int) int {
	return n1 - n2
}

func multiplicar(n1 int, n2 int) (resultado int) {
	resultado = n1 * n2
	return
}

func dividir(n1 int, n2 int) (int, int) {
	quociente := n1 / n2
	resto := n1 % n2
	return quociente, resto
}

func incrementoDecremento(n1 int, n2 int) (inc int, dec int) {
	inc = n1 + n2
	dec = n1 - n2
	return
}
