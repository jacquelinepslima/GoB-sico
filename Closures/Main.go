package main

import (
	"fmt"
)

func main() {
	var numero1, numero2 int
	var operacao string
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanf("%d \n", &numero1)
	fmt.Print("Digite a operação (+, -, *, /, $): ")
	fmt.Scanf("%s \n", &operacao)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%d \n", &numero2)
	var metodoOperacao func(n1 int, n2 int) (int, int)
	if operacao == "+" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 + n2, 0
		}
	} else if operacao == "-" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 - n2, 0
		}
	} else if operacao == "*" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 * n2, 0
		}
	} else if operacao == "/" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 / n2, n1 % n2
		}
	} else if operacao == "$" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 + n2, n1 - n2
		}
	}
	resultado1, resultado2 := metodoOperacao(numero1, numero2)
	fmt.Printf("%d %s %d = %d, %d", numero1, operacao, numero2, resultado1, resultado2)
}
