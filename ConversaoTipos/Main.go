package main

import (
	"fmt"
	"strconv"
)

func main() {
	var texto string
	fmt.Print("Digite um número: ")
	fmt.Scanf("%s", &texto)
	//var numero int
	//numero, _ = strconv.Atoi(texto)
	numero, _ := strconv.ParseInt(texto, 10, 64)
	fmt.Println(numero)
	var conv = float64(numero)
	fmt.Println(conv)

}
