package main

import "fmt"

func main() {
	/* fmt.Println("Inteiros sem sinal")
	var u1 byte = 255 //unit8
	fmt.Println(u1)
	var u2 uint16 = 33
	fmt.Println(u2)
	var u3 uint32 = 789
	fmt.Println(u3)
	var u4 uint64 = 2021
	fmt.Println(u4)
	fmt.Println()

	fmt.Println("Inteiros com sinal")
	var i1 int8 = 127
	fmt.Println(i1)
	var i2 int16 = 33
	fmt.Println(i2)
	var i3 rune = -789 //int32
	fmt.Println(i3)
	var i4 int64 = -2021
	fmt.Println(i4)
	fmt.Println()

	fmt.Println("Interios conforme a arquitetura do processador")
	var t1 uint = 254
	fmt.Println(t1)
	var t2 int = 2000
	fmt.Println(t2)
	fmt.Println()

	fmt.Println("Pontos flutuantes")
	var f1 float32 = 1
	fmt.Println(f1)
	var f2 float64 = 2
	fmt.Println(f2)
	var f3 complex64 = 3
	fmt.Println(f3)
	var f4 complex128 = 4 //suporta os numeros imaginarios + o float64
	fmt.Println(f4)
	fmt.Println()

	fmt.Println("STRINGS")
	var s1 string = "Jacqueline Lima"
	fmt.Println(s1)
	var s2 string = `Hello,
	world` //faz quebra de linha automatica
	fmt.Println(s2)
	fmt.Println(s1[0]) //array
	fmt.Println()

	fmt.Println("BOOLEANOS")
	var b1 bool = true
	fmt.Println(b1)
	var b2 bool = false
	fmt.Println(b2)
	fmt.Println()

	fmt.Println("MULTIPLAS DECLARACOES")
	var nome, sobrenome string = "Jacqueline", "Lima"
	fmt.Println(nome + " " + sobrenome)
	var (
		nome1 string = "TreinaWeb"
		nome2 string = "Cursos"
		idade int    = 10
	) //assim é possivel colocar varios tipos no mesmo lugar
	fmt.Println(nome1 + " " + nome2)
	fmt.Println(idade)

	fmt.Println()

	fmt.Println("Inferencia de tipo")
	var nome = "Jacqueline"
	fmt.Println(nome)
	//em vez de ficar declarando sempre VAR
	//var nomeCompleto string = "Jacqueline Lima"
	nomeCompleto := "Jacqueline Lima"
	fmt.Println(nomeCompleto) */

	fmt.Println()
	fmt.Println("CONSTANTES")
	//const constante string = "Treinaweb Cursos"
	const (
		primeiroNome = "Treinaweb"
		segundoNome  = "Cursos"
	)
	//fmt.Println(constante)
	fmt.Println(primeiroNome + " " + segundoNome)
}
