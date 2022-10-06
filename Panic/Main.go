package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem-vindo!")
	defer fmt.Println("Obrigado por utilizar o nosso software")
	fmt.Print("Digite um número maior que 10: ")
	var numero int
	fmt.Scanf("%d \n", &numero)
	if numero <= 10 {
		panic("Número inválido")
	} else {
		fmt.Println("Você digitou um bom número! :)")
	}
}
