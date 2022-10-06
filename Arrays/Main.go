package main

import (
	"fmt"
)

func main() {
	var amigos [5]string
	fmt.Println(amigos)
	for i := 0; i < len(amigos); i++ {
		fmt.Print("Digite o nome de um dos seus amigos: ")
		fmt.Scanf("%s \n", &amigos[i])
	}
	fmt.Println(amigos)
	fmt.Println("Seus amigos são: ")
	for _, amigo := range amigos {
		fmt.Printf(" - %s \n", amigo)
	}
	arrayInicializado := [3]int{2, 4, 6}
	fmt.Println(arrayInicializado)
	var matriz [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Digite um número: ")
			fmt.Scanf("%d \n", &matriz[i][j])
		}
	}
	fmt.Println(matriz)
}
