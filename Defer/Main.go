package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estou abrindo a conexão com o banco de dados")
	defer fmt.Println("Estou fechando a conexão com o banco de dados!")
	//garante a liberação a recursos, como conexao a banco de dados
	defer fmt.Println("Mais um defer")
	/*quando os defers estao em conjunto sao executados em sistema lifo,
	o ultimo que foi registrado vai ser o primeiro a ser executado*/
	executarConsulta()
}

func executarConsulta() {
	fmt.Println("Estou executando a consulta ao banco de dados...")
}
